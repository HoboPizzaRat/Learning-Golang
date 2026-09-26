package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	jsonData, err := json.Marshal(data)

	// set method to PUT, apply jsonDAta to be PUT
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	// modify request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// make request using http.Client do method
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}

	// decode tand return the response json body(which is user)
	var u User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&u)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)

	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}

	var u User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&u)
	if err != nil {
		return User{}, err
	}

	return u, nil

}
