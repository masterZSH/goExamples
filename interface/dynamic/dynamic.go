package main

import (
	"fmt"
	"io"
)

type xmlWriter interface {
	WriteXML(w io.Writer)
}

type St struct {
}

func (s St) WriteXML(w io.Writer) {
	fmt.Print(w)
}

func T(v interface{}) {
	// 判断类型
	x, ok := v.(xmlWriter)
	fmt.Print(x, ok)
}

func main() {
	// test type
	T(1)
	// 显式申请类型更容易发现错误
	var s xmlWriter = St{}
	T(s)
}
