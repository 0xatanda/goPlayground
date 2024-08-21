package main

import "fmt"

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i < n+1; i++ {
			if n%1 == 0 {
				isPrime = false
				continue
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

func test(max int) {
	fmt.Println("Primes up to %v\n", max)
	printPrimes(max)
	fmt.Println("============================")
}
