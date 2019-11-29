package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.text", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print("open file error")
	}
	file.WriteString("test string\n")
	
}
