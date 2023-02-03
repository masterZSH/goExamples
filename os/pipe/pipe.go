package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write([]byte("123"))
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 1024)
	n, err := r.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data[:n])
	// rescueStdout := os.Stdout
	// r, w, _ := os.Pipe()
	// os.Stdout = w
	// fmt.Println("sample")
	// defer w.Close()
	// out, _ := ioutil.ReadAll(r)
	// os.Stdout = rescueStdout
	// if len(out) == 0 {
	// 	fmt.Println("info log could not be written")
	// }
	// fmt.Printf("read %s\n", out)
}
