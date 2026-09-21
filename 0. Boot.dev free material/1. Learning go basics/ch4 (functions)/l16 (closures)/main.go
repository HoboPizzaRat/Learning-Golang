package main

// A closure is a function that references variables from outside its own function body.
// The function may access and assign to the referenced variables.

func adder() func(int) int {
	sum := 0
	return func(add int) int {
		sum += add
		return sum
	}
}
