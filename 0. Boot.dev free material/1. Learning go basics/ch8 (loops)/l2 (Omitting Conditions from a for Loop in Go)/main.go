package main

import (
	"fmt"
)

//Loops in Go can omit sections of a for loop. For example,
// the CONDITION (middle part) can be omitted which
// causes the loop to run forever.

func maxMessages(thresh int) int {
	var total int = 0
	for i := 0; total <= thresh; i++ {
		if total+100+i < thresh {
			total += 100 + i
		} else {
			return total
		}
	}
	return total
}

func main() {
	fmt.Println(maxMessages(1000))
}
