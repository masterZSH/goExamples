package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

var cwd string

func init() {
	c, err := os.Getwd()
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	cwd = c
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("缺少文件名称\n")
		os.Exit(1)
	}
	fileName := os.Args[1]
	filePath := path.Join(cwd, fileName)
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error : %v", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
