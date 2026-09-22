package main

import (
	"errors"
)

func validateStatus(status string) error {
	if status == "" {
		var err error = errors.New("status cannot be empty")
		return err
	}
	if len(status) > 140 {
		var err error = errors.New("status exceeds 140 characters")
		return err
	}
	return nil
}
