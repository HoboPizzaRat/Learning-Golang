package main

// Go does not support classes. For a long time,
// that meant that Go code couldn't easily be reused in many cases.

// In Go 1.18 however, generics were introduced
/*
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
*/

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	} else {
		return s[len(s)-1]
	}
}
