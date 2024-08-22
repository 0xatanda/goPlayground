package main

import "fmt"

func getFormatedMsg(msgs []string, formatter func(string) string) []string {
	formattedMsg := []string{}
	for _, msg := range msgs {
		formattedMsg = append(formattedMsg, formatter(msg))
	}
	return formattedMsg
}

func addSignature(msg string) string {
	return msg + "Kind regards"
}

func addGreeting(msg string) string {
	return "Hello!!" + msg
}

func tst(msgs []string, formatter func(string) string) {
	defer fmt.Println("=================================")

}

func main() {
	fmt.Println(getFormatedMsg([]string{"bababbababab "}, addSignature))
}
