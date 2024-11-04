package main

import (
	"log"
	"os"
)

var (
	Warn   *log.Logger
	Error  *log.Logger
	Notice *log.Logger
)

func main() {
	warnFile, err := os.OpenFile("warnings.log", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer warnFile.Close()

	Warn = log.New(warnFile, "WARNING: ", log.Ldate|log.Ltime)
	Warn.Println("Messages written to a file called 'warnings.log' are likely to be ignored :(")

	log.Println("Done")
}
