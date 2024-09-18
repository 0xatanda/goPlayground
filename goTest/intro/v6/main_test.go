package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say Hello", func(t *testing.T) {
		got := Hello("Atanda", "")
		want := "Hello, Atanda"
		assertMessage(t, got, want)
	})

	t.Run("Hello, World", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)
	})

	t.Run("Say Hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
