package main

// You can cast an interface to its underlying type using a type assertion.
/*
c, ok := s.(circle)
if ok {
	radius := c.radius
	fmt.Printf("s is a circle, radius: %v\n", radius)
	return
}
*/
func getExpenseReport(e expense) (string, float64) {
	email, ok := e.(email)
	if ok {
		return email.toAddress, email.cost()
	}
	sms, ok := e.(sms)
	if ok {
		return sms.toPhoneNumber, sms.cost()
	}
	return "", 0.0
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

func (em email) cost() float64 {
	if !em.isSubscribed {
		return float64(len(em.body)) * .05
	}
	return float64(len(em.body)) * .01
}

func (sm sms) cost() float64 {
	if !sm.isSubscribed {
		return float64(len(sm.body)) * .1
	}
	return float64(len(sm.body)) * .03
}

func (inv invalid) cost() float64 {
	return 0.0
}
