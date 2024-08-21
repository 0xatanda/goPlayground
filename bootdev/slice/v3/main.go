package main

import (
	"fmt"
)

func sum(nums ...float64) float64 {
	num := 0.0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	return num
}

func test(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("======END OF REPORT ==============")
}

func main() {
	test(1.0, 33.9, 58.3, 0.4, 9.8)
}
