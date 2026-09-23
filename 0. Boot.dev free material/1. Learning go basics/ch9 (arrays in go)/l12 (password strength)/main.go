package main

import "fmt"

// the passowrd must meet the following
// At least 5 characters long but no more than 12 characters.
// Contains at least one uppercase letter.
// Contains at least one digit.

func isValidPassword(password string) bool {
	// ?
	var numberASCIIRange []int = []int{48, 57}
	var upperCaseASCIIRange []int = []int{65, 90}
	var lowerCaseASCIIRange []int = []int{97, 122}

	if len(password) < 5 || len(password) > 12 {
		return false
	}

	var containsNumbers bool = false
	var containsUpperCase bool = false
	var containsLowerCase bool = false

	for i := 0; i < len(password); i++ {
		iAsInteger := int(password[i])
		if iAsInteger >= numberASCIIRange[0] && iAsInteger <= numberASCIIRange[1] {
			containsNumbers = true
		}
		if iAsInteger >= lowerCaseASCIIRange[0] && iAsInteger <= lowerCaseASCIIRange[1] {
			containsLowerCase = true
		}
		if iAsInteger >= upperCaseASCIIRange[0] && iAsInteger <= upperCaseASCIIRange[1] {
			containsUpperCase = true
		}
	}

	if containsLowerCase && containsUpperCase && containsNumbers {
		return true
	}

	return false
}

func main() {
	isValid := isValidPassword("234Huutista")
	fmt.Println(isValid)
}
