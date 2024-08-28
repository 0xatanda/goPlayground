package main

import (
	"errors"
	"fmt"
)

func getLogger(formatter func(string, string) string) func(string, string) {
	return func(s1, s2 string) {
		fmt.Println(formatter(s1, s2))
	}
}

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("===============================================")
	logger := getLogger(formatter)
	fmt.Println("Logs: ")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func coloDelimit(first, second string) string {
	return first + ": " + second
}

func commaDelimit(first, second string) string {
	return first + ", " + second
}

func main() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
	}

	test("Error on database server", dbErrors, coloDelimit)

	mailError := []error{
		errors.New("email to large"),
		errors.New("not enough space to receive mails"),
	}

	test("Error on mail server", mailError, commaDelimit)
}
