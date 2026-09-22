package main

import (
	"fmt"
	"reflect"
)

type contact struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

func main() {
	con := reflect.TypeOf(contact{})
	fmt.Printf("Struct is %d bytes\n", con.Size())

	perm := reflect.TypeOf(perms{})
	fmt.Printf("Struct is %d bytes\n", perm.Size())
}
