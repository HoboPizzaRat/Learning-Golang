package main

// The default case in a select statement executes immediately if no other channel has a value ready.

// ignoring channels
// time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
// time.After() sends a value once after the duration has passed.
// time.Sleep() blocks the current goroutine for the specified duration of time.

// you must add time.Millisecond to your time, otherwise it will be on nanoseconds
/*
time.Tick(500 * time.Millisecond)
*/

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {

	for {
		select {
		case _, ok := <-snapshotTicker:
			if ok {
				takeSnapshot(logChan)
			}
		case _, ok := <-saveAfter:
			if ok {
				saveSnapshot(logChan)
				return
			}
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
