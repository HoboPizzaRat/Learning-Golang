package main

import "fmt"

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func main() {
	e := email{
		message:     "Huutikset",
		fromAddress: "kaurapuuro@gmail.com",
		toAddress:   "kurva@gmail.com",
	}
	e.setMessage("Hiekkalaatikko laatikko laatikko laatikko...")
	fmt.Println(e)
}
