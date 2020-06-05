package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2, str3 string = "test123", "test", "con"
	// true
	fmt.Println(strings.Contains(str1, str2))
	// false
	fmt.Println(strings.Contains(str1, str3))
}
