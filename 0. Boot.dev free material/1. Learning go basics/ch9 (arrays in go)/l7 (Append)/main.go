package main

import "fmt"

// he built-in append function is used to dynamically
// add elements to a slice
/*
slice = append(slice, firstThing, secondThing)
*/
type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	// ?
	slice := []float64{}
	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			slice = append(slice, costs[i].value)
		}
	}
	return slice
}

func main() {
	costs := []cost{
		{
			day:   3,
			value: 3.4475,
		},
		{
			day:   2,
			value: 3.4475,
		},
		{
			day:   2,
			value: 3.4475,
		},
	}
	result := getDayCosts(costs, 2)
	fmt.Println(result)
}
