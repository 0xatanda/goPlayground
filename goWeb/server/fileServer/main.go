package main

import "net/http"

const (
	Port = ":8080"
)

// old way of serving the web
func main() {
	http.ListenAndServe(Port, nil)
	http.FileServer(http.Dir("/var/wwww"))
}
