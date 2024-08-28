package main

import "fmt"

type emailBill struct {
	constInPennies int
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func test(bills []emailBill) {
	defer fmt.Println("========================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've be sent %d emails and it cost you %d cents\n", countAdder(1), costAdder(bill.constInPennies))
	}
}

func main() {
	test([]emailBill{
		{22},
		{76},
		{74},
		{46},
		{22},
	})
}
