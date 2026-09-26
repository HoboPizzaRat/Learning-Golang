package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUsers(url string) ([]User, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Some shit happened")
	}

	var users []User = []User{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&users); err != nil {
		fmt.Println("Error decoding response body")
		return nil, fmt.Errorf("Error decoding response body")
	}
	return users, nil
}
