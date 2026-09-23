package main

//  We can create a new slice using the make function:

// syntax
// mySlice := make([]int, 5, 10)

// in which
// first parameter is the type
// second parameter is the length it contains
// third parameter is the max capacity it will have before needing to resize

import "fmt"

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages), len(messages))
	for i := 0; i < len(messages); i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}

func main() {
	messages := []string{"kkk", "kurvakorva", "kuukkeli"}
	costs := getMessageCosts(messages)
	fmt.Println(costs)
}
