package benchmark

import (
	"math/rand"
	"time"
)

// Int returns a non-negative random number from min to max
// (0 <= min <= nr <= max).
// Will return 0 if max < 0
func Int(min, max int) int {
	var choice int

	if max < 0 {
		return 0
	}
	if min < 0 {
		min = 0
	}

	if min < max {
		choice = min + rand.Intn(max-min+1)
	} else {
		choice = min
	}
	return choice
}

// IntMany returns a slice of non-negative random numbers from min to max.
func IntMany(min, max, quantity int) (out []int) {
	for i := 0; i < quantity; i++ {
		rand.Seed(time.Now().UTC().UnixNano() + int64(i))
		out = append(out, Int(min, max))
	}
	return out
}

// IntSample returns a slice of non-negative random, unique numbers
// from min to max. If the quantity is bigger than max - min sample size
// will be the size of max - min.
func IntSample(min, max, quantity int) (sample []int) {
	if max < 0 {
		return []int{0}
	}
	if min < 0 {
		min = 0
	}
	maxQuantity := max - min
	if maxQuantity > quantity {
		maxQuantity = quantity
	}
	i := 0
	for {
		rand.Seed(time.Now().UTC().UnixNano() + int64(i))
		randNum := Int(min, max)
		if !containsInt(sample, randNum) {
			sample = append(sample, randNum)
			i++
		} else {
			continue
		}
		if i == maxQuantity {
			break
		}
	}
	return sample
}

func containsInt(list []int, elem int) bool {
	for _, t := range list {
		if t == elem {
			return true
		}
	}
	return false
}
