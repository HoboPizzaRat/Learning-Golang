package main

import "fmt"

//A variadic function takes an arbitrary number of final arguments

/*
func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}
// can be called with
final := concat("Hello ", "there ", "friend!")
or
printStrings(final...)
*/

func sum(nums ...int) int {
	var sum int = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...))
}
