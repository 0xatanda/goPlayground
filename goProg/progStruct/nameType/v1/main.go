package main

import "fmt"

type Celsius float64
type Fahrenhelt float64

const (
	AbsoluteZeroC Celsius = -273.15
	Freezing      Celsius = 0
	Boiling       Celsius = 100
)

func CToF(c Celsius) Fahrenhelt {
	return Fahrenhelt(c*9/5 + 32)
}

func FToC(f Fahrenhelt) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C \n", c)
}

func main() {
	var c Celsius
	var f Fahrenhelt

	c = FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v \n", c)
	fmt.Printf("%s \n", c)
	fmt.Println(c)
	fmt.Printf("%g \n", c)
	fmt.Println(float64(c))

	fmt.Println()
	fmt.Println()

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	fmt.Println(c == Celsius(f))
	fmt.Println(f == Fahrenhelt(c))
	fmt.Printf("%g \n", Boiling-Freezing)
	BoilingF := CToF(Boiling)
	fmt.Printf("%g \n", BoilingF-CToF(Freezing))
	fmt.Printf("%g \n", BoilingF-Fahrenhelt(Freezing))
}
