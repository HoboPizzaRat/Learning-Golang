package main

import "fmt"

func main() {
	var age int = 10
	var name string = "kalle"
	var floaty float64 = 2.342433245

	// The %v variant prints any value in a default format.
	s1 := fmt.Sprintf("I am %v years old", age)

	// the %s is for strings
	s2 := fmt.Sprintf("%s is horny", name)

	// the %d si for integers
	s3 := fmt.Sprintf("I amd fucking tarded and %d years old", age)

	// the %f is for floats
	// the %.2f can be used to round number to 2 decimal places
	s4 := fmt.Sprintf("I need to... %.2f", floaty)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
