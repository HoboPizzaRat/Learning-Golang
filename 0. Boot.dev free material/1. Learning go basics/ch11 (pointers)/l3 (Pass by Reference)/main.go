package main

import "fmt"

/*
Functions in Go generally pass variables by value

func increment(x int) {
    x++
    fmt.Println(x)
    // 6
}

func main() {
    x := 5
    increment(x)
    fmt.Println(x)
    // 5
}

One of the most common use cases for pointers in Go is to pass variables by reference
func increment(x *int) {
    *x++
    fmt.Println(*x)
    // 6
}

func main() {
    x := 5
    increment(&x)
    fmt.Println(x)
    // 6
}
*/

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func analyzeMessage(a *Analytics, m *Message) {
	if m.Success == true {
		a.MessagesSucceeded += 1
	} else {
		a.MessagesFailed += 1
	}
	a.MessagesTotal += 1
}
func main() {
	a := Analytics{}
	m := Message{
		Recipient: "Kekoknen",
		Success:   true,
	}
	analyzeMessage(&a, &m)
	fmt.Println(a)
}
