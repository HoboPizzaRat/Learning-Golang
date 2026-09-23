package main

import (
	"fmt"
	"strings"
)

// maps can contain other maps etc
/*
map[string]map[string]int
*/

//Complete the getNameCounts function. It takes a slice of strings names and returns a nested map.
// it returns a following nested map
/*
b: {
    billy: 2,
    bob: 1
},
j: {
    joe: 1
}
*/

func getNameCounts(names []string) map[rune]map[string]int {
	nameCountMap := map[rune]map[string]int{}

	for i := 0; i < len(names); i++ {
		startLetter := []rune(names[i])[0]
		name := strings.ToLower(names[i])

		_, ok := nameCountMap[rune(startLetter)]
		if ok == false {
			nameCountMap[startLetter] = make(map[string]int)
			nameCountMap[startLetter][name] = 1
		} else {
			nameCountMap[startLetter][name]++
		}
	}
	return nameCountMap
}

func main() {
	names := []string{"billy", "billy", "bob", "joe"}
	nameCountMap := getNameCounts(names)
	fmt.Println(nameCountMap)
}
