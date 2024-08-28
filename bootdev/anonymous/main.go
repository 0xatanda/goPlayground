package main

import "fmt"

func printReport(messages []string) {
	for _, message := range messages {
		printCostReport(func(s string) int {
			return len(s) * 2
		}, message)
	}
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf("Messages: %s Cost %d cents\n", message, cost)
}

func test(messages []string) {
	defer fmt.Println("===========================")
	printReport(messages)
}

func main() {
	test([]string{
		"Heres 0x!!!!!!!!!!!!!",
		"Go ahead, make my day",
		"you had me at hello",
	})
	test([]string{
		"Hello, my name is 0x. yu killed my server",
		"show the money!!",
		"Go ahead, make my day",
	})
}
