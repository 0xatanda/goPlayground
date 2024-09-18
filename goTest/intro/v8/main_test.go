package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say Hello to people", func(t *testing.T) {
		got := Hello("Atanda", "English")
		want := "Hello, Atanda"
		assertMessage(t, got, want)
	})

	t.Run("Default Hello World", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertMessage(t, got, want)
	})

	t.Run("Hello in Spanish", func(t *testing.T) {
		got := Hello("Atanda", "Spanish")
		want := "Hola, Atanda"
		assertMessage(t, got, want)
	})

	t.Run("Hello in French", func(t *testing.T) {
		got := Hello("Atanda", "French")
		want := "Bonjour, Atanda"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
