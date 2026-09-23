package main

import "fmt"

// You can combine an if statement with an assignment
// operation to use the variables inside the if block:

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	// ?
	for i := 0; i < len(messagedUsers); i++ {
		username := messagedUsers[i]
		if _, ok := validUsers[username]; ok {
			validUsers[username]++
		}
	}
}
func main() {
	messagedUsers := []string{"kekkonen", "kakkonen", "kukkakauppias", "kekkonen", "kakkonen", "kukkakauppias"}
	validUsers := map[string]int{
		"kekkonen": 0,
	}
	updateCounts(messagedUsers, validUsers)
	fmt.Println(validUsers)
}
