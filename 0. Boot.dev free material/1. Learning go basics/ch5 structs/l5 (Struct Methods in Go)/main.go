package main

// go support methods that can be defined on structs
// methods in this sense are just functions that have
// a receiver

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func main() {
	authInfo := authenticationInfo{
		username: "Kekkonen",
		password: "234823434",
	}
	fmt.Println(authInfo.getBasicAuth())
}
