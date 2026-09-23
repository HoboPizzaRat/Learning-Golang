package main

import (
	"errors"
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(customer *customer, transaction transaction) error {
	switch transaction.transactionType {
	case transactionDeposit:
		customer.balance += transaction.amount
		return nil
	case transactionWithdrawal:
		if customer.balance < transaction.amount {
			return errors.New("insufficient funds")
		} else {
			customer.balance -= transaction.amount
		}
		return nil
	default:
		return errors.New("unkown transaction type")
	}
}
func main() {
	c := customer{
		id:      200,
		balance: 1000,
	}
	t := transaction{
		customerID:      c.id,
		amount:          100,
		transactionType: transactionDeposit,
	}
	updateBalance(&c, t)
	fmt.Println(c)
}
