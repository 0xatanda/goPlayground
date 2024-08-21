package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
	retry1   = "click here to sign up"
	retry2   = "pretty please click here"
	retry3   = "we beg you to sign up"
)

func getMsgPlan(plan string) ([]string, error) {
	allMsg := getMsgRetries()
	if plan == planPro {
		return allMsg[:], nil
	}
	if plan == planFree {
		return allMsg[0:2], nil
	}
	return nil, errors.New("Unsupported plan")
}

func getMsgRetries() [3]string {
	return [3]string{
		retry1,
		retry2,
		retry3,
	}
}

func test(name string, doneAt int, plan string) {
	defer fmt.Println("================================")
	fmt.Printf("Sending to %v....", name)
	fmt.Println()

	msg, err := getMsgPlan(plan)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for i := 0; i < len(msg); i++ {
		msgs := msg[i]
		fmt.Printf(`sending "%v"`, msgs)
		fmt.Println()

		if i == doneAt {
			fmt.Println("they responded")
			break
		}

		if i == len(msg)-1 {
			fmt.Println("no respose")
		}
	}

}

func main() {
	test("0x", 3, planFree)
	test("atanda", 3, planPro)
	test("nafiu", 2, planPro)
	test("omo", 3, "no plan")
}
