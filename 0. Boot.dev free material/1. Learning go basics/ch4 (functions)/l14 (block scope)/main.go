package main

//  Go is not function-scoped, it's block-scoped.

// variables inside any block are only contained within them.
// yea

func splitEmail(email string) (string, string) {
	username, domain := "", ""

	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}
