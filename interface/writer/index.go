package main

import (
	"io"
	"os"
)

func main() {
	var writer io.Writer
	writer = os.Stdout
	io.WriteString(writer)
}
