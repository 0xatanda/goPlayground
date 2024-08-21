package main

import "fmt"

func fizzbuzz() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizzz")
		} else if i%5 == 0 {
			fmt.Println("buzzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
