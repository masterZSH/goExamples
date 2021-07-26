package main

import (
	"bytes"
	"io"
)

func main() {

}

type Buffer struct {
	content *bytes.Buffer
}

func (b *Buffer) Len() int {
	return b.content.Len()
}

func NewBuffer() *Buffer {
	return &Buffer{
		content: bytes.NewBuffer([]byte{}),
	}

}

func (b *Buffer) Write(w io.Writer) (n int, err error) {
	n, err = w.Write(b.content.Bytes())
	return
}
