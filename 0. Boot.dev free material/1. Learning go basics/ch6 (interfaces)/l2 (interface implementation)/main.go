package main

import "fmt"

// mplement the missing getSalary method
// for the contractor type so that it fulfills
// the employee interface.
type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}
func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func main() {
	c := contractor{
		name:         "Kekkonen",
		hourlyPay:    20,
		hoursPerYear: 303,
	}
	fmt.Printf("Name: %s, salary: %d", c.getName(), c.getSalary())
}
