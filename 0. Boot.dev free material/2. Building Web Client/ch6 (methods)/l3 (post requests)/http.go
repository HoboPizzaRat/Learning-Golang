package main

// Like http.Get, the standard library's
// http.Post function can be used to send simple POST requests
import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Take a URL and apiKey string, and User data as parameters
func createUser(url, apiKey string, data User) (User, error) {

	// Encode the user data using json.Marshal
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	// Create a new POST request using http.NewRequest.
	// Use a bytes.NewBuffer to create an io.Reader from the JSON data.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	// Set the Content-Type header, with application/json as its value
	req.Header.Set("Content-Type", "application/json")

	// Set the X-API-Key header, with apiKey as its value
	req.Header.Set("X-API-Key", apiKey)

	// Make the request using the http.Client's Do method
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	// Decode and return the response's JSON body (which is also a User)
	var u User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&u)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
