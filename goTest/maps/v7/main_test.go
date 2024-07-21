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
		definition := "this is just a test"
		dictionary := Dic{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExists)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dic := Dic{word: def}
		newDef := "new definition"
		err := dic.Update(word, newDef)
		assertError(t, err, nil)
		assertDefinition(t, dic, word, newDef)
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
