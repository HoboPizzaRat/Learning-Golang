// these two are the one and the same

// mileage, company := 80276, "Toyota"

// vs

// mileage := 80276
// company := "Toyota"

package main

import "fmt"

func main() {
	averageOpenRate, displayMessage := .23, "is the average open rate of youru messages"

	fmt.Println(averageOpenRate, displayMessage)
}
