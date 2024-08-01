package main

import "fmt"

func main() {
	message := 10
	maxMesage := 20
	fmt.Println("Trying to send a mesage of length:", message, "and a max message of:", maxMesage)

	if message <= maxMesage {
		fmt.Println("mesage sent")
	} else {
		fmt.Println("message not sent")
	}

}
