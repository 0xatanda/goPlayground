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
	resp := "The time is now " + time.Now().String()
	//  Fprintln allows direct output to any writer
	fmt.Fprintln(w, resp)
}

// serveStatic func allow to serve any filedirectly and use Go as an old-sch
// web server that serves only static content
func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func serveError(w http.ResponseWriter, r *http.Request) {
	fmt.Println("There's no way this will work")
}

func main() {
	http.HandleFunc("/static", serveStatic)
	http.HandleFunc("/", serveDynamic)
	http.HandleFunc("/error", serveError)
	http.ListenAndServe(Port, nil)
}
