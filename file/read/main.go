package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	c, err := Content("f")
	fmt.Println(err)
	fmt.Printf("%s\n", c)
}

func Content(f string) (content []byte, err error) {
	fi, err := os.Open(f)
	if err != nil {
		return
	}
	defer fi.Close()
	buf := make([]byte, 10)
	for {
		n, err := fi.Read(buf)
		content = append(content, buf[:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return content, err
		}

	}
	return content, nil
}
