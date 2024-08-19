package main

import "fmt"

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that cost $%v to be sent to %v can not be sent", cost, recipient)
}

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("===========================")
}

func main() {
	test(1.4, "+234 (903) 864 8068")
	test(44.3, "234 (807) 142 0648")
}
