package main

import "fmt"

func main() {
	// p type of *int, piont to an unamed int variable
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}
