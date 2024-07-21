package main

import (
	"bytes"
	"fmt"
)

func Great(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
