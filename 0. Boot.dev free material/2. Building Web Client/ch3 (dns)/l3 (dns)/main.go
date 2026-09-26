package main

import (
	"fmt"
	"net/url"
)

// The net/url package is part of Go's standard library.
// You can instantiate a URL struct using url.Parse:
/*
parsedURL, err := url.Parse("https://homestarrunner.com/toons")
if err != nil {
	fmt.Println("error parsing url:", err)
	return
}
*/
func getDomainNameFromURL(rawURL string) (string, error) {
	parseURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("error parsing url", err)
		return "", err
	}
	return parseURL.Hostname(), nil
}
