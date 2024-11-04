package test

import "testing"

func TestSquare(t *testing.T) {
	if v := Square(4); v != 16 {
		t.Error("Expected", 16, "got", v)
	}
}
