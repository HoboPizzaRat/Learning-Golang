package main

// 99 times out of 100 you will use a slice instead
// of an array when working with ordered lists.

/*
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}
*/

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	// ?
	if plan == planPro {
		return messages[:], nil
	} else if plan == planFree {
		return messages[:2], nil
	} else {
		return nil, errors.New("unsupported plan")
	}
}

func main() {
	messages := [3]string{"kekkonen", "kakkonen", "kukkakauppias"}
	result, err := getMessageWithRetriesForPlan(planFree, messages)
	if err != nil {
		fmt.Println("Huutista")
	}
	fmt.Println(result)
}
