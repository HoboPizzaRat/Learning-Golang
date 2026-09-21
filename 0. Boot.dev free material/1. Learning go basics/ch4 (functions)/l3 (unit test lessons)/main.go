//
// Complete the getMonthlyPrice function.
// It accepts a tier (string) as input and returns the monthly price for that tier in pennies.
// Here are the prices in dollars:

//    "basic" - $100.00
//    "premium" - $150.00
//    "enterprise" - $500.00

package main

func getMonthlyPrice(tier string) int {
	if tier == "basic" {
		return 100_00
	} else if tier == "premium" {
		return 150_00
	} else if tier == "enterprise" {
		return 500_00
	}
	return 0
}
