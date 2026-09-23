package main

// Methods with pointer receivers can modify the value to which the receiver points

// methods with pointer receivers don't require that a pointer is used to call the method
import (
	"fmt"
)

type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
}
