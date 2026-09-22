package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func main() {
	user1 := user{
		name:   "Kekkonen",
		number: 1234,
	}
	user2 := user{
		name:   "Kakkonen",
		number: 3223,
	}
	mToSend := messageToSend{
		message:   "Heyy! Wanna have a match?",
		sender:    user1,
		recipient: user2,
	}
	var result bool = canSendMessage(mToSend)
	fmt.Println(result)
}
