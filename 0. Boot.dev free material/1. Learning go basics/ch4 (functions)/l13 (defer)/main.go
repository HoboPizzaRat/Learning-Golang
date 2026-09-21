package main

// The defer keyword is a fairly unique feature of Go.
// It allows a function to be executed automatically
// just before its enclosing function returns.

// Deferred functions are typically used to clean up resources that
// are no longer being used. Often to close
// database connections, file handlers and the like.

/* example
func GetUsername(dstName, srcName string) (username string, err error) {
	// Open a connection to a database
	conn, _ := db.Open(srcName)

	// Close the connection *anywhere* the GetUsername function returns
	defer conn.Close()

	username, err = db.FetchUser()
	if err != nil {
		// The defer statement is auto-executed if we return here
		return "", err
	}

	// The defer statement is auto-executed if we return here
	return username, nil
}
*/
import "fmt"

func bootup() {
	// task
	// be sure to rpint "TEXTIO BOOTUP DONE" before the bootup function returns
	// use defer so that you ahve to write this msg once instead
	// before each return statement
	ok := connectToDB()
	defer fmt.Println("TEXTIO BOOTUP DONE")
	if !ok {
		return
	}
	ok = connectToPaymentProvider()
	if !ok {
		return
	}
	fmt.Println("All systems ready!")
}

var shouldConnectToDB = true

func connectToDB() bool {
	fmt.Println("Connecting to database...")
	if shouldConnectToDB {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

var shouldConnectToPaymentProvider = true

func connectToPaymentProvider() bool {
	fmt.Println("Connecting to payment provider...")
	if shouldConnectToPaymentProvider {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

func test(dbSuccess, paymentSuccess bool) {
	shouldConnectToDB = dbSuccess
	shouldConnectToPaymentProvider = paymentSuccess
	bootup()
	fmt.Println("====================================")
}

func main() {
	test(true, true)
	test(false, true)
	test(true, false)
	test(false, false)
}
