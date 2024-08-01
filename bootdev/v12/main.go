package main

import "fmt"

func main() {
	sendSoFar := 430
	const sendToAdd = 25
	sendSoFar = increment(sendSoFar, sendToAdd)
	fmt.Println("you've sent", sendSoFar, "messages")
}

func increment(sendSoFar, sendToAdd int) int {
	sendSoFar += sendToAdd
	return sendSoFar
}
