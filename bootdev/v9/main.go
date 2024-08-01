package main

import "fmt"

func main() {
	const name = "0x"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is %f percent", name, openRate)

	fmt.Println(msg)
}
