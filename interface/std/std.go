package main

import (
	"bufio"
	"fmt"
	"os"
)

// io interface case
func main() {
	fmt.Fprintf(os.Stdout, "test")
	// use bufio
	bufWriter := bufio.NewWriter(os.Stdout)
	bufWriter.WriteString("io test")
	defer bufWriter.Flush()
}
