package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	Port = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// mux.vars looks for string query from the http.Request
	// and parse them into maps.
	vars := mux.Vars(r)

	// The value can be accessible by referencing the result by key,
	// in this case "id"
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	http.ServeFile(w, r, fileName)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	rtr.HandleFunc("/homepage", pageHandler)
	rtr.HandleFunc("/cntact", pageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(Port, nil)
}
