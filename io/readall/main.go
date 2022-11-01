package main

import (
	"fmt"
	"io"
)

func main() {

	// ReadAll Added in 1.16

	r, w := io.Pipe()
	go func() {
		w.Write([]byte("12345"))
		w.Close()
	}()
	result, err := io.ReadAll(r)
	fmt.Printf("%s\n,%+v", result, err)
}
