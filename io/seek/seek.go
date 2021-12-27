package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("t")
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, f)

	r, err := f.Seek(10, io.SeekStart)
	fmt.Println(r, err)
}
