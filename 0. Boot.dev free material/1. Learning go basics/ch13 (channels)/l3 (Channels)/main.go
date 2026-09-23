package main

// in previus lesson we saw how we can receive values from channels like this:
// v := <-ch

// sometimes we don't care what is passed through a channel.
// <-ch

// In cases like this, empty structs are often used as a unary value so that the sender means its only a "signal"

import "fmt"

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}
