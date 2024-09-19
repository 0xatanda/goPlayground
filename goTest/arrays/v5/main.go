package main

func Sum(nums []int) int {
	sums := 0
	for _, num := range nums {
		sums += num
	}
	return sums
}

func SumAll(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}
