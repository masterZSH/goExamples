package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fmt.Println("sample")
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if len(out) == 0 {
		fmt.Println("info log could not be written")
	}
	fmt.Printf("read %s\n", out)
}
