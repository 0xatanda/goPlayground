package main

import "fmt"

func yearsUtilEvent(age int) (yearsUtilAdult, yearsUtilDrinking, yearsUtilCarRental int) {
	yearsUtilAdult = 18 - age
	if yearsUtilAdult < 0 {
		yearsUtilAdult = 0
	}

	yearsUtilDrinking = 21 - age
	if yearsUtilDrinking < 0 {
		yearsUtilDrinking = 0
	}

	yearsUtilCarRental = 25 - age
	if yearsUtilCarRental < 0 {
		yearsUtilCarRental = 0
	}

	return yearsUtilAdult, yearsUtilDrinking, yearsUtilCarRental
}

func main() {
	fmt.Println(yearsUtilEvent(20))
}
