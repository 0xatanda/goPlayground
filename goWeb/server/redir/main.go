package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const Port = ":8080"

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	err, _ := os.Stat(fileName)
	if err != nil {
		fileName = "files/404.html"
	}
	http.ServeFile(w, r, fileName)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(Port, nil)
}