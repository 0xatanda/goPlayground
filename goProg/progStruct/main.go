package main

import "fmt"

func main() {
	const boilingF, freezingF = 212.0, 32.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToc(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToc(boilingF))
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
