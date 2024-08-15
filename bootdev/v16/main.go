package main

import "fmt"

func yrUnit(age int) (yrUtilAdult int, yrUtilDrinking int, yrUtilCar int) {
	yrUtilAdult = 18 - age
	if yrUtilAdult < 0 {
		yrUtilAdult = 0
	}

	yrUtilDrinking = 21 - age
	if yrUtilDrinking < 0 {
		yrUtilDrinking = 0
	}

	yrUtilCar = 25 - age
	if yrUtilCar < 0 {
		yrUtilCar = 0
	}

	return
}

func main() {
	fmt.Println(yrUnit(30))
}
