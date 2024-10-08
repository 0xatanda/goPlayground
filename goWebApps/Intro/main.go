package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	Port = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	res := "This time is now " + time.Now().String()
	fmt.Fprintln(w, res)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func main() {

	http.HandleFunc("/", serveDynamic)

	http.HandleFunc("/static", serveStatic)
	fmt.Println("server starting on port 8080")
	http.ListenAndServe(Port, nil)
}
