package main

// Go has a language feature, that when used with a type definition
// (and if you squint really hard), kinda looks like
// an enum (but it's not). It's called iota

type emailStatus int

const (
	EmailBounced emailStatus = iota
	Emailnvalid
	EmailDelivered
	EmailOpened
)
