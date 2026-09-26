package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

/*
sometimes you need to deal with json data of unknown structures

var data map[string]interface{}
jsonString := `{"name": "Alice", "age": 30, "address": {"city": "Wonderland"}}`
json.Unmarshal([]byte(jsonString), &data)
fmt.Println(data["name"])  // Output: Alice
fmt.Println(data["address"].(map[string]interface{})["city"])  // Output: Wonderland
*/

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &resources); err != nil {
		return nil, err
	}
	return resources, nil

}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	for _, resource := range resources {
		for key, value := range resource {
			var str string = fmt.Sprintf("Key: %v - Value: %v\n", key, value)
			formattedStrings = append(formattedStrings, str)
		}
	}
	sort.Strings(formattedStrings)

	fmt.Println()
	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}
