package main

import (
	"fmt"
)

// refrences
// https://golang.org/doc/effective_go.html#interfaces_and_types

// Stringer 字符串接口
type Stringer interface {
	String() string
}

// Interface conversions and type assertions
func main() {
	var value interface{} // Value provided by caller.
	value = "123"
	fmt.Printf("%s", GetStr(value))
}

func GetStr(v interface{}) string {
	switch str := v.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}
	return ""
}
