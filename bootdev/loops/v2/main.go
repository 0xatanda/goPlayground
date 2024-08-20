package main

import "fmt"

func maxMsg(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (1.0 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

func test(thresh float64) {
	fmt.Println("Threshold: %.2f\n", thresh)
	max := maxMsg(thresh)
	fmt.Printf("Maxium msg that can be sent: = %v\n", max)
}
