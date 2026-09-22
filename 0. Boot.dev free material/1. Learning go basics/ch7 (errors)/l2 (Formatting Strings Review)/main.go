package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%s' cannot be sent", cost, recipient)
}

func main() {
	fmt.Println(getSMSErrorString(03.4, "kekoknen on kakonen"))
}
