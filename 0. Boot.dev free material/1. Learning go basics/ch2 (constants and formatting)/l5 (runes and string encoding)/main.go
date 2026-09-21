package main

import (
	"fmt"
	"unicode/utf8"
)

// utf.RuneCountInString() can be used to count the total characters of string

// normal len function only returns the amount of bytes the string has
// in go you can have unicode characters in your string so it breaks the normal len
func main() {
	const name = "boots"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}
