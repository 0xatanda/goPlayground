package main

import (
	"fmt"
	"net/http"
	"time"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintln(w, t.YearDay())
}

func main() {
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}
