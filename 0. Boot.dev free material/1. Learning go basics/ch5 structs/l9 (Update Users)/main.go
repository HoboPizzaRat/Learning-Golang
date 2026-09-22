package main

import "fmt"

// Create a new struct called Membership, it should have:
// A Type string field
// A MessageCharLimit integer field
type Membership struct {
	Type             string
	MessageCharLimit int
}

// Update the User struct to embed a Membership.
type User struct {
	Name       string
	membership Membership
}

// Complete the newUser function.
// It should return a new User with all the fields set as you
// would expect based on the inputs. If the user is a "premium" member,
// the MessageCharLimit should be 1000, otherwise, it should only be 100.
func newUser(name string, membershipType string) User {
	// ?
	var user User
	if membershipType == "premium" {
		user = User{
			Name: name,
			membership: Membership{
				Type:             "premium",
				MessageCharLimit: 1000,
			},
		}
	} else {
		user = User{
			Name: name,
			membership: Membership{
				Type:             "normal",
				MessageCharLimit: 100,
			},
		}
	}
	return user
}
func main() {
	var user User = newUser("Kekkonen", "premium")
	fmt.Println(user)
}
