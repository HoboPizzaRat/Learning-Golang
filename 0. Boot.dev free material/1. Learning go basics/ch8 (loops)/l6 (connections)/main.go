package main

import "fmt"

func countConnections(groupSize int) int {
	var connections int = 0
	for i := groupSize - 1; i > 0; i-- {
		connections += i
	}
	return connections
}

func main() {
	fmt.Println(countConnections(1000))
}
