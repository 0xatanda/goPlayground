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

	errorFile, err := os.OpenFile("error.log", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer errorFile.Close()

	Warn = log.New(warnFile, "WARNING: ", log.LstdFlags)

	Warn.Println("Messages written to a file called 'warnings.log' are likely to be ignored :(")
	Error = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime)
	Error.SetOutput(errorFile)
	Error.Println("Error messages, on the other hand, tend to catch attention!")
}
