package main

import (
	"errors"
	"fmt"
)

/*
insert element
m[key] = elem

get an element
elem = m[key]

delete an element
delete(m, key)

// check if key exists
elem, ok := m[key]
*/

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	_, ok := users[name]
	if ok == false {
		return false, errors.New("not found")
	}
	if users[name].scheduledForDeletion == false {
		return false, nil
	}
	delete(users, name)
	return true, nil

}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {
	deleteMap := map[string]user{
		"kekkonen": {
			name:                 "kekkonen",
			number:               28349034,
			scheduledForDeletion: true,
		},
		"kukkakauppias": {
			name:                 "kukkakauppias",
			number:               28349034,
			scheduledForDeletion: true,
		},
		"kakkonen": {
			name:                 "kakkonen",
			number:               28349034,
			scheduledForDeletion: true,
		},
	}
	result, err := deleteIfNecessary(deleteMap, "kekkonen")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(deleteMap)
}
