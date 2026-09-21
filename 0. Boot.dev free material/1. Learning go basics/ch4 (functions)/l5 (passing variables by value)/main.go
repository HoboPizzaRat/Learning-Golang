// Variables in Go are passed by value (except for a few data types we haven't covered yet)

// assignment

// monthlyBillIncrease:
// Should return the increase in the bill from the previous to
// the current month. If the bill decreased, return a negative number.

// getBillForMonth:
// Should return the total cost for the number of messages sent

package main

import "fmt"

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	lastMonthBill = getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill = getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	bill := costPerSend * messagesSent
	return bill
}

func main() {
	costPerSend := 10
	numLastMonth := 20
	numThisMonth := 35
	fmt.Println(monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth))
}
