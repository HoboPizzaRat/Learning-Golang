package main

import (
	"encoding/json"
	"log"
)

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TaskCount   int    `json:"taskCount"`
}

func marshalAll[T any](items []T) ([][]byte, error) {

	marshaled := make([][]byte, len(items))

	for i := 0; i < len(items); i++ {
		data, err := json.Marshal(items[i])
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		marshaled[i] = make([]byte, len(data))
		marshaled[i] = data
	}
	return marshaled, nil
}
