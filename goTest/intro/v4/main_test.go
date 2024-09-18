package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Atanda")
	want := "Hello, Atanda"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
