package main

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	lenOfNums := len(numsToSum)
	sum := make([]int, lenOfNums)

	for i, num := range numsToSum {
		sum[i] = Sum(num)
	}
	return sum
}
