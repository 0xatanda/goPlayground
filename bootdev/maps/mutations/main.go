package main

import (
	"errors"
	"fmt"
)

type user struct {
	name              string
	number            int
	scheduledDeletion bool
}

func deleteIfNecessary(user map[string]user, name string) (deleted bool, err error) {
	existingUser, ok := user[name]
	if !ok {
		return false, errors.New("not found")
	} else if existingUser.scheduledDeletion {
		delete(user, name)
		return true, nil
	}
	return false, nil
}

func test(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s....\n", name)
	defer fmt.Println("===============================")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("Deleted: ", name)
		return
	}
	fmt.Println("Did not delete: ", name)
}

func main() {
	users := map[string]user{
		"0x": {
			name:              "0x",
			number:            9957573832,
			scheduledDeletion: true,
		},
	}

	test(users, "0x")
}
