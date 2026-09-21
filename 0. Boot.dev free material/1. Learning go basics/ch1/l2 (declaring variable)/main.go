package main

import "fmt"

func main() {
	// we can declare variables with type with var

	// take a note that the variable name and the type are backwards
	// esp compared  to other languages such as c or java
	var mySkillIssues int

	// the prorgram wont compile if we dont use our declared variables
	// in the program by design
	mySkillIssues = 42

	// we print shit
	// we dont necessarily need to convert int types to string to print them
	fmt.Println(mySkillIssues)
}
