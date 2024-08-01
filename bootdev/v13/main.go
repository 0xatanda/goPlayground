package main

import "fmt"

func main() {
	firstName, _ := getName()
	fmt.Println("Welcome to Lagos", firstName)
}

func getName() (string, string) {
	return "0x", "atanda"
}
