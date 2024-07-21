package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dic := Dic{"test": "this is just a test"}
	dict := Dic{}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("Unknown Word", func(t *testing.T) {
		_, got := dic.Search("unknown")
		assertError(t, got, ErrNotFound)
	})

	t.Run("Add Words", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict.Add(word, def)
		assertDefinition(t, dict, word, def)
	})

	t.Run("New Word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		err := dict.Add(word, def)
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a just"
		dic := Dic{word: def}
		err := dict.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, def)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dic Dic, word, def string) {
	t.Helper()

	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertString(t, got, def)
}
