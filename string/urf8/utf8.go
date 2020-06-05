package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string = "test测试"
 	num:= utf8.RuneCountInString(str)
	fmt.Println(num)        // 6
	fmt.Println(len(str))   // 4+10中文长度3
}