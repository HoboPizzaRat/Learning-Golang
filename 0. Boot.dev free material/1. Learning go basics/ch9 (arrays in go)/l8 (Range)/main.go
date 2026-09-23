package main

import "fmt"

// This is the go way to iterate through all items from a list

/*
fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
*/

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func main() {
	msg := []string{"Kakka", "kekkonen", "kukkakauppias"}
	badWords := []string{"kukkakauppias", "hitto", "vittu"}

	var index int = indexOfFirstBadWord(msg, badWords)
	fmt.Println(index)
}
