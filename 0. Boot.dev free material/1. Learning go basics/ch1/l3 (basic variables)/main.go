package main

import "fmt"

// these are the most common variables of go

func main() {
	// whole integer values are denoted by int
	var health int
	health = 1000

	// boolean values are dnoe with bool keyword
	// the true false optionsn are lowercase
	var amIRetarded bool
	amIRetarded = true

	// string values are done with string
	var swearWord string
	swearWord = "Voi perkele"

	// floating points are done with float64
	// there are also float32 bit but use 64-bit preferrably
	var floatingPoint float64
	floatingPoint = 3.234234

	// we can also handle data on a byte level and there are names just for that.
	var data byte
	data = 0xaa

	fmt.Println(health)
	fmt.Println(amIRetarded)
	fmt.Println(swearWord)
	fmt.Println(floatingPoint)
	fmt.Println(data)
}
