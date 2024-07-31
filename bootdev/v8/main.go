package main

import "fmt"

func main() {
	const secondInMinutes = 60
	const minutesInHour = 60
	const secondInHour = secondInMinutes * minutesInHour

	fmt.Println("number of seconds in an hour", secondInHour)
}
