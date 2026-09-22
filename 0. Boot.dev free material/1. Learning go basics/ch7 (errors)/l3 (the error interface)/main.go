package main

import (
	"fmt"
)

// this is how you can build a custom error interface

/*
	type userError struct {
	    name string
	}

	func (e userError) Error() string {
	    return fmt.Sprintf("%v has a problem with their account", e.name)
	}
*/
type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("cannot divide %f by zero", d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
