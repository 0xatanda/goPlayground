package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Great(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreat(w http.ResponseWriter, r *http.Request) {
	Great(w, "0x")
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(MyGreat)))
}
