package main

import "fmt"

type messageToSend struct {
	msg   string
	phone uint
}

func test(m messageToSend) {
	fmt.Printf("sending msg: '%s' to %v\n", m.msg, m.phone)
	fmt.Println("========================================")
}

func main() {
	test(messageToSend{
		phone: 9038877209,
		msg:   "Thanks for signing up",
	})
}
