package main

/*
Two strings can be concatenated with the + operator.
But the compiler will not allow you to concatenate
a string variable with an int or a float64
*/

import "fmt"

func main() {
	var msg string = "some weird thingy from: "
	var user string = "Kekkonen"
	var output string = msg + user

	fmt.Println(output)
}
