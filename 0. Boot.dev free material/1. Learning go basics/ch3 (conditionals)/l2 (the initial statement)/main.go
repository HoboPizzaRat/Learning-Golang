package main

/*
the if condition can have 'initial' statement
it can create the statement used for the condition

if INITIAL_STATEMENT; CONDITION {
	do some conditions in here
}
*/
import "fmt"

// practical example
func main() {
	var email string = "kekkonen@gmail.com"

	if length := len(email); length < 10 {
		fmt.Printf("Email must be at least 10 characters long!")
	}
}
