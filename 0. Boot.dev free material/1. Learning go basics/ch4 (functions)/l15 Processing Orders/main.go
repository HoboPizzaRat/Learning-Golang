package main

// complete the placeholder functio
// it returns a bool indicating whether the order was successfull

// The placeOrder function should always return the
// account balance regardless of whether it was adjusted.

// the amountInStock and calcPrice functions can be
// used to look up the current stock
// and price of the item

// The amountInStock and calcPrice functions can be
// used to look up the current stock and price of an item.
func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
	if quantity > amountInStock(productID) {
		return false, accountBalance
	}
	if accountBalance < calcPrice(productID, quantity) {
		return false, accountBalance
	}
	accountBalance -= calcPrice(productID, quantity)
	return true, accountBalance
}

func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}

func priceList(productID string) float64 {
	if productID == "1" {
		return 1.50
	} else if productID == "2" {
		return 2.25
	} else if productID == "3" {
		return 3.00
	} else if productID == "4" {
		return 1.00
	} else if productID == "5" {
		return 2.50
	} else if productID == "6" {
		return 8.99
	} else if productID == "7" {
		return 22.50
	} else if productID == "8" {
		return 50.00
	} else if productID == "9" {
		return 999.99
	} else {
		return 0.00
	}
}

func amountInStock(productID string) int {
	if productID == "1" {
		return 11
	} else if productID == "2" {
		return 25
	} else if productID == "3" {
		return 4
	} else if productID == "4" {
		return 6
	} else if productID == "5" {
		return 50
	} else if productID == "6" {
		return 2
	} else if productID == "7" {
		return 0
	} else if productID == "8" {
		return 99
	} else if productID == "9" {
		return 1
	} else {
		return 0
	}
}
