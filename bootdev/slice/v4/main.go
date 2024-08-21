package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostByDay(costs []cost) []float64 {
	costByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		if cost.day >= len(costByDay) {
			costByDay = append(costByDay, 0.0)
		}
		costByDay[cost.day] += cost.value
	}
	return costByDay
}

func test(cost []cost) {
	fmt.Printf("Creating daily bucket for %v costs\n", len(cost))
	costsByDay := getCostByDay(cost)
	fmt.Println("cost by Days: ")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("=======END REPORT===============================")
}

func main() {
	test([]cost{
		{0, 1.0},
		{1, 2.0},
		{2, 5.5},
	})
}
