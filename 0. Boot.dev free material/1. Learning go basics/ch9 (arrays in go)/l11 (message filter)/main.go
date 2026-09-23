package main

import "fmt"

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

func filterMessages(messages []Message, filterType string) []Message {

	var filtered []Message = []Message{}

	for i := 0; i < len(messages); i++ {

		if messages[i].Type() == filterType {
			filtered = append(filtered, messages[i])
		}
	}
	return filtered
}

func main() {
	var messages []Message = []Message{
		TextMessage{
			Sender:  "Juutalainen",
			Content: "salajuonia...",
		},
		LinkMessage{
			Sender:  "Kekkonen",
			URL:     "mp4",
			Content: "Huutista",
		},
		MediaMessage{
			Sender:    "Kekkonen",
			MediaType: "https://keekoslandia.com",
			Content:   "kekkonen metsällä...",
		},
	}

	media := filterMessages(messages, "media")
	fmt.Println(media)
}
