package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {
	// data写入到 w.data
	// func Write(w io.Writer, order ByteOrder, data interface{}) error
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, []byte("test data"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("% x\n", buf.Bytes())

	// bs 创建Buffer字节缓冲区
	bs := new(bytes.Buffer)
	binary.Write(bs, binary.LittleEndian, uint32(14))
	fmt.Printf("% x\n", bs.Bytes())
}
