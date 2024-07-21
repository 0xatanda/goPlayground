package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct {
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

const finalWord = "Go!"
const Countdownstart = 3

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleep Sleeper) {
	for i := Countdownstart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleep.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
