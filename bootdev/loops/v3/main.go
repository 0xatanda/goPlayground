package main

import "fmt"

func getMaxMsg(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMsgToSend := 0
	for actualCostInPennies <= float64(maxCostInPennies) {
		maxMsgToSend++
		actualCostInPennies *= costMultiplier
	}
	return maxMsgToSend
}

func test(costMultiplier float32, maxCostInPennies int) {
	fmt.Printf("Cost %2.f and max cost: %v \n", costMultiplier, maxCostInPennies)
	cost := getMaxMsg(float64(costMultiplier), maxCostInPennies)
	fmt.Println("The cost is %v", cost)
}
