package main

import (
	"fmt"
	"time"
)

// Interfaces are just collections of method signatures.
// A type "implements" an interface if it has methods
// that match the interface's method signatures.

func sendMessage(msg message) (string, int) {
	return msg.getMessage(), len(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}
func main() {
	msg := birthdayMessage{
		birthdayTime:  time.Now(),
		recipientName: "Kekkonen",
	}
	sentMessage, length := sendMessage(msg)
	fmt.Println(sentMessage)
	fmt.Println(length)
}
