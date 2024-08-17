package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	receipent user
}

type user struct {
	name   string
	number int
}

func canSend(mToS messageToSend) bool {
	if mToS.sender.name == "" {
		return false
	}
	if mToS.receipent.name == "" {
		return false
	}
	if mToS.sender.number == 0 {
		return false
	}

	if mToS.receipent.number == 0 {
		return false
	}

	return true
}

func test(mToS messageToSend) {
	fmt.Printf("sending '%s' from %s (%v) to %s (%v).....sent!\n", mToS.message, mToS.sender.name, mToS.sender.number, mToS.receipent.name, mToS.receipent.number)
}

func main() {
	test(messageToSend{
		message: "you have an appointment tomorrow",
		receipent: user{
			name:   "atanda nafiu",
			number: 9033666883,
		},
		sender: user{
			name:   "atanda kolapo",
			number: 9129484558,
		},
	})
}
