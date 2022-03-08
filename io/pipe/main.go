package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// r, w := io.Pipe()
	// // defer w.Close()
	// go io.Copy(os.Stdout, r)
	// go func() {
	// 	w.Write([]byte("123"))
	// }()
	// for {
	// }

	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
