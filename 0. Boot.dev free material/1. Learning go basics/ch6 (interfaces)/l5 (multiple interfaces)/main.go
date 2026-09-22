package main

import (
	"fmt"
)

// somethign can implement multipel interfaces duch

func (e email) cost() int {
	costPerChar := 0
	if e.isSubscribed {
		costPerChar = 2
	} else {
		costPerChar = 5
	}
	return costPerChar * len(e.body)
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("'%s' | Subscribed", e.body)
	} else {
		return fmt.Sprintf("'%s' | Not Subscribed", e.body)
	}
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

func main() {
	e := email{
		isSubscribed: false,
		body:         "Kekkonen on kultakakkukana",
	}
	fmt.Println(e.format())
	fmt.Println(e.cost())

}
