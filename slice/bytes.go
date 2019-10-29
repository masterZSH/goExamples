package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 读写长度未知的bytes最好使用Buffer结构体
	var r *bytes.Buffer = new(bytes.Buffer)
	fmt.Printf("%p", r)
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok { //method getNextString() not shown here
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Print(buffer.String(), "\n")
}

func getNextString() (string, bool) {
	return "234", false
}
