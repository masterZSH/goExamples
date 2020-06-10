package main

import (
	"io"
	"os"
)

func main() {
	bf := []byte("test")
	var w io.Writer
	// w.Write(bf) // nil pointer
	w = os.Stdout
	w.Write(bf)
}

//
