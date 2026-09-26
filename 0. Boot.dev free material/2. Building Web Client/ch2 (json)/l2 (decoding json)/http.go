package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// To decode JSON into a slice of Issue structs,
// we need to know the JSON fields and their types.

/*
type Issue struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Estimate int    `json:"estimate"`
}
*/

// after receiving a json response, we can decode it into a slice of
// of "issue" structs with address of opeartor (&)
/*
var issues []Issue
decoder := json.NewDecoder(res.Body)
if err := decoder.Decode(&issues); err != nil {
	fmt.Println("error decoding response body")
	return
}
for _, issue := range issues {
    fmt.Printf("Issue – id: %v, title: %v, estimate: %v\n", issue.Id, issue.Title, issue.Estimate)
    // Issue – id: 001-a, title: Unspaghettify code, estimate: 9001
}
*/

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue = []Issue{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&issues); err != nil {
		fmt.Println("Error decoding response body")
		return nil, fmt.Errorf("Error decoding response body")
	}
	return issues, nil
}
