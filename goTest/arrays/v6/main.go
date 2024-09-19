package main

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAllTail(numsToSum ...[]int) []int {
	var sum []int
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sum = append(sum, 0)
		} else {
			tail := nums[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return sum
}
