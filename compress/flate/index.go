package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func main() {
	w := new(bytes.Buffer)
	// ignore error
	writer, _ := flate.NewWriter(w, flate.BestCompression)
	writer.Write([]byte("ces1231"))
	writer.Flush()
	defer writer.Close()
	fmt.Println(w)

	// 创建一个读取器
	deCompressor := flate.NewReader(w)
	defer deCompressor.Close()
	// 输出到标准输出
	io.Copy(os.Stdout, deCompressor)
}
