package main

// You can provide a buffer length as the second
// argument to make() to create a buffered channel:
// ch := make(chan int, 100)

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, 1000)
	for i := 0; i < len(emails); i++ {
		ch <- emails[i]
	}
	return ch
}

func main() {
	emails := []string{"Kekkonen", "jumalauta", "hiekkalaatikko"}
	addEmailsToQueue(emails)
}
