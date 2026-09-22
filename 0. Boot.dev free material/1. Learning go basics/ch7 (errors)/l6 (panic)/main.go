package main

// The panic function yeets control out of the current function and up
// the call stack until it reaches a function that defers a recover.
// If no function calls recover, the goroutine
// (often the entire program) crashes.

/*
func enrichUser(userID string) User {
    user, err := getUser(userID)
    if err != nil {
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    // this panics, but the defer/recover block catches it
    // a truly astonishingly bad way to handle errors
    enrichUser("123")
}
*/
