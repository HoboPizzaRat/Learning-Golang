package main

/*
When multiple arguments are of the same type, and are next to each other in the function signature,
the type only needs to be declared after the last argument.
*/

// here are some examples
func addToDatabase(hp, damage int) {
	// ...
}
func addToDatabase(hp, damage int, name string) {
	// ?
}
func addToDatabase(hp, damage int, name string, level int) {
	// ?
}
func main() {
	addToDatabase()
	// etc etc
}
