package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

func test(s sender) {
	fmt.Println("sender name:", s.name)
	fmt.Println("sender number: ", s.number)
	fmt.Println("send rateLimit: ", s.rateLimit)
}

func main() {
	test(sender{
		rateLimit: 10,
		user: user{
			name:   "atanda nafiu",
			number: 9904857,
		},
	})
}
