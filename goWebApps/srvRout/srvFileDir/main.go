package main

import "net/http"

const (
	Port = ":8080"
)

func main() {
	http.ListenAndServe(Port, http.FileServer(http.Dir("/var/www")))
}
