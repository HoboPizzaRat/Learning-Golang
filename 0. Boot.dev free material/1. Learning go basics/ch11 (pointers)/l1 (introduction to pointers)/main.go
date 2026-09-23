package main

// pointers are eh varibales (ariwhales) that stores the memory address of data
// and it can manipulate the data its pointing to

/*
var p *int

myString := "hello"
myStringPtr := &myString
*/

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, m.Recipient, m.Text)
}

func main() {
	message := Message{
		Recipient: "Kekkonen",
		Text:      "Huutista",
	}
	result := getMessageText(message)
	fmt.Println(result)
}
