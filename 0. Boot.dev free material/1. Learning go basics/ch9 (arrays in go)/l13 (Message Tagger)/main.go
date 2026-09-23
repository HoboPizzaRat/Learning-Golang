package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

// this function takes slice of sms messgaes, and returns a sms
// list containing the same sms messgaes but they have been
// given tag properties given by tagger function
func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// ?
	var messagesTagged []sms = []sms{}
	for i := 0; i < len(messages); i++ {
		tagged := sms{
			id:      messages[i].id,
			content: messages[i].content,
			tags:    tagger(messages[i]),
		}
		messagesTagged = append(messagesTagged, tagged)
	}
	return messagesTagged
}

// tagger function takes single sms message in and
// returns the message tagged if it contains
// "urgent", then tag added is Urgent
// "sale", then tag added is Promo
func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "urgent")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}

func main() {
	var messages []sms = []sms{
		{id: "001", content: "Urgent! Last chance to see!"},
		{id: "002", content: "Big sale on all items!"},
	}
	var tagged []sms = tagMessages(messages, tagger)

	for i := 0; i < len(tagged); i++ {
		fmt.Println(tagged[i])
	}
}
