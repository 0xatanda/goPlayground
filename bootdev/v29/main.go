package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

func test(x, y float64) {
	defer fmt.Println("=========================")
	fmt.Printf("Dividing %.2f by %.2f ..\n", x, y)
	quotient, err := divide(x, y)
	if err != nil {
		errors.New("Cant be divided")
	}
	fmt.Println(quotient)
}

func main() {

}
