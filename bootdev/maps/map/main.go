package main

import (
	"errors"
	"fmt"
)

type Usr struct {
	name        string
	phoneNumber int
}

func getUser(names []string, phoneNumbers []int) (map[string]Usr, error) {
	userMap := make(map[string]Usr)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]

		userMap[name] = Usr{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

func test(names []string, phoneNumber []int) {
	fmt.Println("Creating maps.....")
	defer fmt.Println("============================")
	users, err := getUser(names, phoneNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println("- name:", users[name].name)
		fmt.Println("- numbers:", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"0x", "0m0", "jill"},
		[]int{123344222, 4472525171, 990381827},
	)
}
