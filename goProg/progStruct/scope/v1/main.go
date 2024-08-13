package main

import "fmt"

func f()  {}

var g = "g"

func main() {
	f := "f"
	fmt.Println(f)  // "f" local var f shadow pkg-level func f
	fmt.Println(g)  // "g" pkg-level var
}