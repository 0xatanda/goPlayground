package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say Hello", func(t *testing.T) {
		got := Hello("Atanda")
		want := "Hello, Atanda"

		if got != want {
			t.Errorf("got %q but want  %q", got, want)
		}
	})

	t.Run("Hello, World", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}
