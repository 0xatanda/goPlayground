package main

import "fmt"

func Rverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func main() {
	fmt.Println(Rverse("Hello"))
}
