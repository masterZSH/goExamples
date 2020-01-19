package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFileWithSlice("test.text")
}

func openFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return
	}
	defer file.Close()
	iReader := bufio.NewReader(file)
	for {
		str, err := iReader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Printf("input %s", str)
	}
}
