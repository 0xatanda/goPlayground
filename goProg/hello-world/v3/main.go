package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 prints the text of each line that appear more than
// once in the standard input, preceded by its count

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// Note: ignore potential errors from Input.err()
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
