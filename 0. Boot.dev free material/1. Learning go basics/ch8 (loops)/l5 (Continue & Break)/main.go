package main

// The continue keyword stops the current iteration of a loop and continues to the next iteration

// The break keyword stops the current iteration of a loop and exits the loop.

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		var isPrime bool = true
		for i := 3; i <= int(math.Sqrt(float64(n))+1); i += 2 {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
