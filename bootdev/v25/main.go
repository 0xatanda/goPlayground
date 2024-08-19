package main

import "fmt"

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type expense interface {
	cost() float64
}

type invalid struct{}

func getExpenseReort(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func estimateYrCt(e expense, avarageMsgYr int) float64 {
	return e.cost() * float64(avarageMsgYr)
}

func test(e expense) {
	address, cost := getExpenseReort(e)
	switch e.(type) {
	case email:
		fmt.Printf("Repot: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("============================================================")
	case sms:
		fmt.Printf("Report: the sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("=============================================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("============================================================")
	}

}

func main() {
	test(email{
		isSubscribed: true,
		body:         "Hello there",
		toAddress:    "0x@gmail.com",
	})
	test(sms{
		isSubscribed:  true,
		body:          "Hello!!!!!!!!!!!!!!!!!!!!",
		toPhoneNumber: "909337565",
	})
}
