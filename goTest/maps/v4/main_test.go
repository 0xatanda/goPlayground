package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dic := Dic{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dic.Search("unknown")
		assertError(t, got, ErrNotFound)
	})

	t.Run("Add words", func(t *testing.T) {
		dic := Dic{}
		dic.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dic.Search("test")
		assertError(t, err, nil)
		assertString(t, got, want)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
