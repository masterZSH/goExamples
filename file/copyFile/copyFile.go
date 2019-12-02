package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	copyFile("copy", "test.text")
	copyIo("copy.gif", "1.gif")
}

func copyFile(dstName, srcName string) (res bool, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return false, err
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	defer dst.Close()
	if err != nil {
		return false, err
	}
	_, err = io.Copy(dst, src)
	if err == nil {
		return true, nil
	}
	return false, nil
}

func copyIo(dstName, srcName string) {
	content, err := ioutil.ReadFile(srcName)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s", content)
	err = ioutil.WriteFile(dstName, content, 0644)
	if err != nil {
		panic(err.Error())
	}
}
