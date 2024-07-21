package main

import "testing"

func TestSearch(t *testing.T) {
	dic := map[string]string{"test": "this is just a test"}

	got := Search(dic, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q but want %q, given %q", got, want, "test")
	}
}
