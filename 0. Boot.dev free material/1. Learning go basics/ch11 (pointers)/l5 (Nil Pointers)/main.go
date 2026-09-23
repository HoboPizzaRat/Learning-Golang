package main

// Pointers can be very dangerous.

// if points doesnt point to anything (is nil)
// deferencing it will cause a panic during runtime

// you should always check pointer before dereferencing

import (
	"strings"
)

func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")

	if message == nil {
		return
	}
	*message = messageVal
}
