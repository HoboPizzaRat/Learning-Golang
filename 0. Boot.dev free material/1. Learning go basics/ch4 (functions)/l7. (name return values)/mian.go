package main

import "fmt"

func yearsUntilEvents(age int) (int, int, int) {
	// don't touch below this line
	var yearsUntilAdult int = 0
	var yearsUntilDrinking int = 0
	var yearsUntilCarRental int = 0

	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

func main() {
	_, drinking, _ := yearsUntilEvents(21)
	fmt.Println(drinking)
}
