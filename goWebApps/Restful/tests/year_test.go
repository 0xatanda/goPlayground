package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	res := httptest.NewRecorder()
	path := "http://localhost:4000/test"
	req, _ := http.NewRequest("GET", path, nil)
	http.DefaultServeMux.ServeHTTP(res, req)
	response, err := io.ReadAll(res.Body)
	if string(response) != "115" || err != nil {
		t.Errorf("Expected [], got %s", string(response))
	}
}
