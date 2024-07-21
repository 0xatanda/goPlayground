package main

import (
	"fmt"
	"io"
	"os"
)

func Great(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Great(os.Stdout, "Atanda0x")
}
