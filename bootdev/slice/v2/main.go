package main

import "fmt"

func getMsgCost(msg []string) []float64 {
	costs := make([]float64, len(msg))
	for i := 0; i < len(msg); i++ {
		message := msg[i]
		cost := float64(len(message))
		costs[i] = cost
	}
	return costs
}

func test(msg []string) {
	costs := getMsgCost(msg)
	fmt.Println("Msg: ")
	for i := 0; i < len(msg); i++ {
		fmt.Printf("-%v\n", msg[i])
	}
	fmt.Println("Costs: ")
	for i := 0; i < len(costs); i++ {
		fmt.Printf("cost : %v\n", costs)
	}
}
