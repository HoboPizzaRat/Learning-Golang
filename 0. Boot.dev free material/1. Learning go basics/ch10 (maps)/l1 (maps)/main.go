package main

import (
	"errors"
	"fmt"
)

// We can create a map by using the make() function:
/*
ages := make(map[string]int)
ages["John"] = 37
*/

//Map values can be structs too:
/*
type car struct {
  registration string
  model        string
}

cars := map[string]car{}
cars["XYZ-789"] = car{registration: "XYZ-789", model: "Accord"}
*/
func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	userMap := make(map[string]user)

	for i := 0; i < len(names); i++ {
		u := user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
		userMap[names[i]] = u
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func main() {
	names := []string{"kekkonen", "kakkonen", "kukkakauppias"}
	phoneNumbers := []int{23478342, 3874384, 7348872394}

	userMap, err := getUserMap(names, phoneNumbers)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userMap)
}
