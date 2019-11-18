package main

import (
	"fmt"
	"unsafe"
)

// File file struct
type File struct {
	fd   int    // 描述符
	name string // 文件名称
}

// NewFile create File
func NewFile(fd int, name string) *File {
	if fd > 0 {
		return nil
	}
	return &File{fd, name}
}

func main() {
	//
	var f = NewFile(1, "231222222222")
	var size = unsafe.Sizeof(&f)
	fmt.Print(size)
}
