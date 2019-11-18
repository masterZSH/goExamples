package main

import (
	"fmt"
	"reflect"
)

// File file struct
type File struct {
	fd   int    // 描述符
	name string // 文件名称
}

func main() {
	var f = &File{1, "231222222222"}
	var types = reflect.TypeOf(f)
	fmt.Printf("%+v\n", types)
}
