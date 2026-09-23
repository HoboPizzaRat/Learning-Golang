package main

// It's possible to define an empty pointer. For example, an empty pointer to an integer:
/*
var p *int

fmt.Printf("value of p: %v\n", p)
// value of p: <nil>
*/

// Instead of starting with a nil pointer, it's common to use the & operator to get a pointer to its operand
/*
myStringPtr := &myString
*/

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	// ?
	mask := strings.Repeat("*", len(*message))
	*message = strings.ReplaceAll(*message, *message, mask)
}
func main() {
	message := "Kekkonen"
	removeProfanity(&message)
	fmt.Println(message)
}
