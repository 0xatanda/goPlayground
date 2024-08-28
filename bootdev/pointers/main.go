package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "***********")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "***")
	*message = messageVal
}

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	test([]string{
		"0x",
	})
}
