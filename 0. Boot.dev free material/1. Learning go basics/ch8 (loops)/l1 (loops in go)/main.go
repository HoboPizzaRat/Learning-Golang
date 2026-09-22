package main

import "fmt"

// The basic loop in Go is written in standard C-like syntax:
/*
for INITIAL; CONDITION; AFTER{
  // do something
}
*/

func bulkSend(numMessages int) float64 {
	// ?
	var totalCost float64 = 0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + 0.01*float64(i)
	}
	return totalCost
}

func main() {
	fmt.Println(bulkSend(3))
}
