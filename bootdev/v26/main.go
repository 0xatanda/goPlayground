package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSoupse string) (float64, error) {
	costFC, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, nil
	}
	costFS, err := sendSMS(msgToSoupse)
	if err != nil {
		return 0.0, nil
	}
	return costFC + costFS, nil
}

func sendSMS(msg string) (float64, error) {
	const mexTextLen = 25
	const costPerchar = .002
	if len(msg) > mexTextLen {
		return 0.0, fmt.Errorf("cn't send text over %v character", mexTextLen)
	}
	return costPerchar * float64(len(msg)), nil
}

func test(msgToCustomer, msgToSoupse string) {
	defer fmt.Println("======================")
	fmt.Println("Message for Customer: ", msgToCustomer)
	fmt.Println("Message for Spouse: ", msgToSoupse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSoupse)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func main() {

}
