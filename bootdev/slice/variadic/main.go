package main

import "fmt"

func sun(nums ...int) int {
	// num is just a slice
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]

	}
	return num
}

func main() {
	total := sun(1, 2, 3, 4, 5)
	fmt.Println(total)
}
