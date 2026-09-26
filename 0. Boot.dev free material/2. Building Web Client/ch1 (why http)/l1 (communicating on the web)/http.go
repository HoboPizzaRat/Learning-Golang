package main

import (
	"fmt"
	"io"
	"net/http"
)

func getIssueData() ([]byte, error) {
	// we can get data from endpoint with using http modules get method
	res, err := http.Get("https://api.boot.dev/v1/courses_rest_api/learn-http/issues")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	return data, err
}
