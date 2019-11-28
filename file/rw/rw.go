package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("./file/test.dat")
	if err != nil {
		fmt.Print(err)
		return
	}
	// 关闭
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		input, readErr := inputReader.ReadString('\n')
		fmt.Println(input)
		if readErr == io.EOF {
			return
		}
	}
}
