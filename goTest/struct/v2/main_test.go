package main

import (
	"testing"
)

func TestAreas(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()
		if got != want {
			t.Errorf("got %g but want %g", got, want)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rectangles := Rectange{12, 6}
		checkArea(t, rectangles, 72.0)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
