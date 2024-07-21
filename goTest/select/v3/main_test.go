package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	slowSever := makeDelay(20 * time.Millisecond)
	fastServer := makeDelay(0 * time.Millisecond)

	defer slowSever.Close()
	defer fastServer.Close()

	slowURL = slowSever.URL
	fastURL = fastServer.URL

	got := fastURL
	want := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func makeDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
