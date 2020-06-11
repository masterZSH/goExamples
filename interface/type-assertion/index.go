package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var i interface{}
	// 空接口类型断言

	i = os.Stdout
	f := i.(*os.File)
	// panic: interface conversion: interface {} is *os.File, not *bytes.Buffer
	w := i.(*bytes.Buffer)
	fmt.Println(f)

	fmt.Println(w)

}
