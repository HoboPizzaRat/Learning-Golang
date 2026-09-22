package main

import "fmt"

// A type switch is similar to a regular switch statement,
// but the cases specify types instead of values.

/*
func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}
*/

func getExpenseReport(e expense) (string, float64) {
	// ?
	switch exp := e.(type) {
	case email:
		return exp.toAddress, exp.cost()
	case sms:
		return exp.toPhoneNumber, exp.cost()
	default:
		return "", 0.0
	}
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "kekkonen on kakkonen",
		toAddress:    "emt@gmail.com",
	}
	s := sms{
		isSubscribed:  true,
		body:          "Kukkanen on kekkonen",
		toPhoneNumber: "32378497",
	}
	fmt.Println(getExpenseReport(e))
	fmt.Println(getExpenseReport(s))

}
