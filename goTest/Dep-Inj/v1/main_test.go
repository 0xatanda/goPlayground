package main

import (
	"bytes"
	"testing"
)

func TestGreat(t *testing.T) {
	buffer := bytes.Buffer{}
	Great(&buffer, "Atanda0x")

	got := buffer.String()
	want := "Hello, Atanda0x"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
