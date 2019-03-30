package main

import (
	"fmt"

	benchmark "github.com/multiverse-os/benchmark"
)

func main() {
	fmt.Println("Benchmarking Empty String Testing")
	fmt.Println("=================================")

	benchmark.TestBenchmark()

	//result := benchmark.RunTest(100, func() { FunctionB("alice") })
	//result.Print()

	// Determine best way to test empty string}
	comparison := benchmark.CompareFunctions(25000, func() { MethodA() }, func() { MethodB() }, func() { MethodC() })
	bestResult := comparison.Best()
	bestResult.Print()
	fmt.Println("Best function is:", comparison.BestFunction)

}

func MethodA() {
	if benchmark.RandomString(1, 20) == "" {
		fmt.Println("empty; this should not be possible")
	}
}

func MethodB() {
	if len(benchmark.RandomString(1, 20)) == 0 {
		fmt.Println("empty; this should not be possible")
	}
}

var emptyByte byte

func MethodC() {
	if benchmark.RandomString(1, 20)[0] == emptyByte {
		fmt.Println("empty; this should not be possible")
	}
}
