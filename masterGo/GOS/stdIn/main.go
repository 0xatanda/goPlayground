package main

import (
	"bufio"
	"fmt"
	"os"
)

var f *os.File

func main() {
	f = os.Stdin
	defer f.Close()

	fmt.Println("============== Write on the terminal ===================")
	fmt.Println()
	fmt.Println()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

}
