package main

func Sum(num [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += num[i]
	}
	return sum
}
