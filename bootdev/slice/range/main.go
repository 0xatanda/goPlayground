package main

import "fmt"

func indexWord(msg, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func test(msg []string, badWords []string) {
	i := indexWord(msg, badWords)
	fmt.Printf("Scanning msg: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(`-`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("=============================================")
}

func main() {
	badWords := []string{"0x", "shoot", "dang", "frick"}
	msg := []string{"hey", "there", "john"}
	test(msg, badWords)
}
