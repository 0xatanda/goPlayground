package main

import (
	"fmt"
	"time"
)

type message interface {
	getMessage() string
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

type birthdayMsg struct {
	birthdayTm    time.Time
	receipentName string
}

func sendMsg(msg message) {
	fmt.Println(msg.getMessage())
}

func (bm birthdayMsg) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.receipentName, bm.birthdayTm)
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`your "%s" report is ready %v`, sr.reportName, sr.reportName)
}

func test(m message) {
	sendMsg(m)
	fmt.Println("=======================")

}

func main() {
	test(sendingReport{
		reportName:    "first report",
		numberOfSends: 10,
	})
	test(birthdayMsg{
		receipentName: "0x",
		birthdayTm:    time.Date(2004, 10, 18, 0, 0, 0, 0, time.UTC),
	})
}
