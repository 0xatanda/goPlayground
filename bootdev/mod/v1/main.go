package main

import "fmt"

func modOp(x, y int) int {
	return x % y
}

func main() {
	fmt.Println(modOp(12, 4))
	fmt.Println(modOp(16, 5))
	fmt.Println(modOp(22, 8))
	fmt.Println(modOp(27, 4))

}
