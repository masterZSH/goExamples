package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	// 修改字符串的一个字符
	var s1 string = "1234567"
	fmt.Println(changeChar(s1, 1, 's'))

	// 获取字符串字串
	fmt.Println(s1[2:4])

	// 获取字符数
	fmt.Println(len(s1))

	// 获取字符数 最快速
	fmt.Println(utf8.RuneCountInString(s1))

	// 连接字符串
	var buff bytes.Buffer
	buff.WriteString(s1)
	buff.WriteString("aaaaaa")
	fmt.Print(buff.String())
}

func changeChar(str string, idx int, char byte) string {
	if len(str) < idx {
		panic("error")
	}
	c := bytes.NewBufferString(str)
	bytes := c.Bytes()
	bytes[idx] = char
	return string(bytes)
}
