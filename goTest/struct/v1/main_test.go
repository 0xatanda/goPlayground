package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Parameters", func(t *testing.T) {
		rectangle := Rectange{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		AssertMsg(t, got, want)
	})

	t.Run("Areas", func(t *testing.T) {
		rectangle := Rectange{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		AssertMsg(t, got, want)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g but want %g", got, want)
		}
	})
}

func AssertMsg(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}
