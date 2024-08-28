package main

import "fmt"

const (
	logDelete  = "user deleted"
	logNotFund = "user not found"
	logAdmin   = "admin deleted"
)

type user struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(users, name)
	user, ok := users[name]
	if !ok {
		return logNotFund
	}
	if user.admin {
		return logAdmin
	}

	return logDelete
}

func test(user map[string]user, name string) {
	fmt.Printf("Attempting to delete %s.....\n", name)
	defer fmt.Println("===============================")
	log := logAndDelete(user, name)
	fmt.Println("log: ", log)
}
