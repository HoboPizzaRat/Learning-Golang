package main

// go programs express errors with error values
//
// When something can go wrong with a function,
// that should return an error as its last return value

/*
type error interface {
    Error() string
}
*/

// atoi
// we can convert string into number
/*
i, err := func Atoi(s string)
if err != nil{
	// handle errors gracefully
}
*/
import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	// ?
	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		fmt.Println("error has occured")
		return 0, err
	}
	fmt.Println(cost)
	fmt.Println("SMS has been sent successfully")

	cost2, err2 := sendSMS(msgToSpouse)
	if err2 != nil {
		fmt.Println("error has occured")
		return 0, err2
	}
	fmt.Println(cost2)
	fmt.Println("SMS has been sent successfully")

	return cost + cost2, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func main() {
	total, err := sendSMSToCouple("kekkonen on kakkonen", "kukkakauppias")
	if err != nil {
		fmt.Println("error sending messages")
		return
	}
	fmt.Println(total)
}
