package main

import "testing"

func TestSearch(t *testing.T) {
	dic := map[string]string{"test": "this is just a test"}
	got := Dic.Search(dic, "test")
	want := "this is just a test"
	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
