/*
he size (8, 16, 32, 64, 128, etc.) represents how many bits in memory
will be used to store the variable. The "default" int and uint types refer
to their respective 32 or 64-bit sizes depending on the environment of the user.
*/

// signed integers
// int  int8  int16  int32  int64

// unsigned integers
// uint uint8 uint16 uint32 uint64 uintptr

// signed decimals
// float32 float64

// complex numbers
// complex64 complex128package main
package main

import "fmt"

// converting between types
func main() {
	temperatureFloat := 88.26
	temperatureInt := int(temperatureFloat)
	fmt.Println(temperatureInt)
}
