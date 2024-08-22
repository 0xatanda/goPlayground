package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func aggregate(a, b, c int, aritmetic func(int, int) int) int {
	return aritmetic(aritmetic(a, b), c)
}

func main() {
	fmt.Println(aggregate(2, 3, 5, add))
	fmt.Println(aggregate(4, 3, 2, mul))
}
