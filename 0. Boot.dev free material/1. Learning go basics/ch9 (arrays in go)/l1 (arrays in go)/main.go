package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}

	var retries [3]string
	var costs [3]int
	var totalCost = 0

	for i := 0; i < 3; i++ {
		retries[i] = messages[i]
		costs[i] = len(messages[i]) + totalCost
		totalCost = costs[i]
	}
	return retries, costs
}

func main() {
	primary := "kekkonen"
	secondary := "kekkonen"
	tertiary := "kekkonen"

	retries, costs := getMessageWithRetries(primary, secondary, tertiary)

	fmt.Println(retries)
	fmt.Println(costs)
}
