package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowURL := "http://www.facebook.com"
		fastURL := "http://www.quii.dev"

		slowSever := makeDelay(20 * time.Millisecond)
		fastServer := makeDelay(0 * time.Millisecond)

		defer slowSever.Close()
		defer fastServer.Close()

		slowURL = slowSever.URL
		fastURL = fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})

	t.Run("Returns an error if a server doesn't response within 10s", func(t *testing.T) {
		serverA := makeDelay(11 * time.Second)
		serverB := makeDelay(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
