package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("please give me one or ore floats")
		os.Exit(1)
	}

	arguments := os.Args
	var err error = errors.New("an error")
	k := 1
	var n float64

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!!")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	min, max := n, k
	for i := 2; i < len(arguments); i++ {
		a, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if a < min {
				min = a
			}
			if a > float64(max) {
				max = int(a)
			}
		}
	}

	fmt.Println("min: ", min)
	fmt.Println("max: ", max)
}
