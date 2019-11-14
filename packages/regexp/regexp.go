package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 正则匹配练习
	pat := "[0-9]+"
	t := "stta"
	ok, _ := regexp.Match(pat, []byte(t))
	fmt.Print(ok)
}
