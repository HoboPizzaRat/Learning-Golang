package main

import "fmt"

// constants can be computed as long as the commputation
// happen at compile time

func main() {
	const firstName = "Lane"
	const lastName = "Wagner"
	const fullName = firstName + " " + lastName
	fmt.Println(fullName)
}
