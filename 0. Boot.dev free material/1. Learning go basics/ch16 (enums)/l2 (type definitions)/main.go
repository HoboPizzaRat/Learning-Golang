package main

type sendingChannel string

const (
	Email sendingChannel = "email"
	SMS   sendingChannel = "sms"
	Phone sendingChannel = "phone"
)

func sendNotification(ch sendingChannel, message string) {
	// send the message
}
