package main

import "fmt"

const (
	retry1 = "click here to sign up"
	retry2 = "pretty please click here"
	retry3 = "we beg you to sign up"
)

func getMsg() [3]string {
	return [3]string{
		retry1,
		retry2,
		retry3,
	}
}

func test(name string, doneAt int) {
	fmt.Printf("Sending to %v....", name)
	fmt.Println()

	msg := getMsg()
	for i := 0; i < len(msg); i++ {
		msgs := msg
		fmt.Printf(`sending "%v"`, msgs)
		fmt.Println()

		if i == doneAt {
			fmt.Println("they responded")
			break
		}
		if i == len(msg)-1 {
			fmt.Println("complete failure")
		}
	}
}

func main() {
	test("0x", 0)
	test("atanda", 1)
	test("nafiu", 2)
	test("omo", 3)
}
