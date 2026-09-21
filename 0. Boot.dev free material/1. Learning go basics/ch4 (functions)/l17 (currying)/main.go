package main

import (
	"errors"
	"fmt"
)

// Function currying is a concept from functional programming and involves
// partial application of functions. It allows a function with multiple arguments
// to be transformed into a sequence of functions, each taking a single argument.

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a string, b string) {
		fmt.Println(formatter(a, b))
	}
}

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}

func commaDelimit(first, second string) string {
	return first + ", " + second
}

func main() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test("Error on mail server", mailErrors, commaDelimit)
}
