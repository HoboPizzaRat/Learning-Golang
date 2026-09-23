package main

import "fmt"

/*
// Similar to slices and maps, channels can be ranged over.
// the range iteration will close when the close signal has been received

for item := range ch {
    // item is the next value received from the channel
}
*/

func concurrentFib(n int) []int {
	result := []int{}
	ch := make(chan int)

	go fibonacci(n, ch)
	for item := range ch {
		result = append(result, item)
	}
	return result
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	result := concurrentFib(10)
	fmt.Println(result)
}
