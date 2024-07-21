package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const Countdownstart = 3

func Countdown(out io.Writer) {
	for i := Countdownstart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
