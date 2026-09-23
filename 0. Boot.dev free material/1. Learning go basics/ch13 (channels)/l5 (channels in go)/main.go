package main

import "fmt"

// Channels can be explicitly closed by a sender:
/*
ch := make(chan int)
// do some stuff with the channel
close(ch)
*/

// Sending on a closed channel will cause a panic!

func countReports(numSentCh chan int) int {
	count := 0
	for {
		_, ok := <-numSentCh
		if ok == false {
			break
		} else {
			fmt.Println("report received!")
			count++
		}
	}
	return count
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

func main() {
	channel := make(chan int)
	numBatches := 10
	go sendReports(numBatches, channel)
	go countReports(channel)
}
