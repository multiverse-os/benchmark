package benchmark

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

const debug = false

func Start(startedAt time.Time) time.Duration {
	elapsed := time.Since(startedAt)

	fmt.Printf("[benchmark] test took [ %s ms ] to complete \n", elapsed)
	return elapsed
}

//////////////////////////////////////////////////////////////////////////////
func TestBenchmark() {
	defer Start(time.Now())
	r := rand.Intn(10)
	time.Sleep(time.Duration(5+r) * time.Microsecond)
}

///////////////////////////////////////////////////////////////////////////////
type Comparison struct {
	Results      []Result
	BestFunction func()
}

func (self Comparison) Best() (result Result) {
	for i, testResult := range self.Results {
		if i == 0 || testResult.AverageRuntime() < result.AverageRuntime() {
			result = testResult
		}
	}
	return result
}

type Result struct {
	Runtimes []time.Duration
}

func (self Result) Print() {
	if debug {
		fmt.Println("[debug] result:", self)
	}
	fmt.Println("[benchmark] test has ran [", len(self.Runtimes), "times ] with an average of [", self.AverageRuntime(), " microseconds ]")
}

func (self Result) AverageRuntime() time.Duration {
	var total time.Duration
	for _, runtime := range self.Runtimes {
		total += runtime
	}
	if len(self.Runtimes) > 0 {
		averageRuntime := total / (time.Duration(len(self.Runtimes)) * time.Nanosecond)
		return averageRuntime
	} else {
		return 0
	}
}

func (self Result) AppendRuntime(runtime time.Duration) Result {
	self.Runtimes = append(self.Runtimes, runtime)
	return self
}

func RunTest(iterations int, testFunction func()) (result Result) {
	for i := 1; i <= iterations; i++ {
		result = result.AppendRuntime(Test(testFunction))
	}
	if debug {
		result.Print()
	}
	return result
}

func CompareFunctions(iterations int, testFunctions ...func()) (comparison Comparison) {
	for _, function := range testFunctions {
		result := RunTest(iterations, function)
		fmt.Println("[benchmark] function being tested [", function, "] with the name [", runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name(), "]")
		result.Print()
		comparison.Results = append(comparison.Results, result)
		comparison.BestFunction = function
	}
	return comparison
}

func Test(testFunction func()) time.Duration {
	startTime := time.Now()
	testFunction()
	return time.Since(startTime)
}
