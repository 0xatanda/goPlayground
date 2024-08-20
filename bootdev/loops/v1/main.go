package main

import "fmt"

func bulkSend(numMessage int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessage; i++ {
		totalCost += 1.0 + (0.1 * float64(i))
	}
	return totalCost
}

func test(numMessage int) {
	fmt.Printf("Sending %v msg\n", numMessage)
	cost := bulkSend(numMessage)
	fmt.Println("Bulk send complete! cost = %.2f\n", cost)
	fmt.Println("========================================")
}

func main() {
	test(int(bulkSend()))
}
