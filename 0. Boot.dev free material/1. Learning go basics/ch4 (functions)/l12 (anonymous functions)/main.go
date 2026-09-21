package main

// Anonymous functions are true to form in that they have no name.

// this is how you define an anonymous function
/*
// some code
func(a int) int {
	    return a + a
}
// some code
*/

// assignment

import "fmt"

func printReports(intro, body, outro string) {
	printCostReport(func(data string) int {
		return len(data) * 2
	}, intro)
	printCostReport(func(data string) int {
		return len(data) * 3
	}, body)
	printCostReport(func(data string) int {
		return len(data) * 4
	}, outro)
}

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
