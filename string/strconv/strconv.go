package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int => string
	fmt.Println(strconv.Itoa(123))
	// 转为二进制
	fmt.Println(strconv.FormatInt(123, 2))
	fmt.Printf("%b\n", 123)
	
}
