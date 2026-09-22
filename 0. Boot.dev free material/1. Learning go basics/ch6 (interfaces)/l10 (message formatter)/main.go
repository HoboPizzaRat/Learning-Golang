package main

import "fmt"

type formatter interface {
	format() string
}

type plainText struct {
	message string
}
type bold struct {
	message string
}
type code struct {
	message string
}

func (p plainText) format() string {
	return fmt.Sprintf("%s", p.message)
}
func (b bold) format() string {
	return fmt.Sprintf("**%s**", b.message)
}
func (c code) format() string {
	return fmt.Sprintf("`%s`", c.message)
}

func sendMessage(format formatter) string {
	return format.format()
}

func main() {
	p := plainText{
		message: "Kukkakaupppias",
	}
	b := bold{
		message: "Kukkakaupppias",
	}
	c := code{
		message: "Kukkakaupppias",
	}
	fmt.Println(sendMessage(p))
	fmt.Println(sendMessage(b))
	fmt.Println(sendMessage(c))
}
