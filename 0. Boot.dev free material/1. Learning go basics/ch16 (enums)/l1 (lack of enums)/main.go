package main

import "errors"

// My least favorite part of Go?
// It's Go's lack of enums, sum types, tagged unions, etc.

func (a *analytics) handleEmailBounce(em email) error {
	err := em.recipient.updateStatus(em.status)
	if err != nil {
		return errors.New("error udpating user status: " + err.Error())
	}
	err2 := a.track(em.status)
	if err2 != nil {
		return errors.New("error tracking user bounce: " + err2.Error())
	}
	return nil
}
