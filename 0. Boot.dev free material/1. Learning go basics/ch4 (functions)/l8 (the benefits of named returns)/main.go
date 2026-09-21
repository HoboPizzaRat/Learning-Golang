package main

// Named return parameters are great for documenting a function.
// We know what the function is returning directly
// from its signature, no need for a comment.

// named return parameters are very useful on longer functions
// with complex return value schemas

import "errors"

// do things like this:
func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

// not like this
/*
func calculator(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}
	mul := a * b
	div := a / b
	return mul, div, nil
}
*/
