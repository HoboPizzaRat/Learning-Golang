package main

/*
messages := []string{"Hello world", "hello there", "General Kenobi"}

count := countDistinctWords(messages)
*/

// create countDistinctWords function using a map
// it shoudl take a slice of strings as arguments
// returns total count of distinct words on map
import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	distinctCount := 0
	distinctMap := make(map[string]int)

	for x := 0; x < len(messages); x++ {
		words := strings.Fields(strings.ToLower(messages[x]))

		for i := 0; i < len(words); i++ {
			word := words[i]
			_, ok := distinctMap[word]

			if ok == false {
				distinctCount++
				distinctMap[word] = 1
			} else {
				distinctMap[word]++
			}
		}
	}
	fmt.Println(distinctMap)
	return distinctCount
}

func main() {
	messages := []string{"kekkonen on kukkakauppias", "kakkonen kelloseppä se on", "huutikset kekkonen"}
	count := countDistinctWords(messages)
	fmt.Println(count)
}
