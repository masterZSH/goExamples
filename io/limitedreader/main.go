package main

import (
	"bytes"
	"fmt"
	"io"
)

type LimitedReader struct {
	R io.Reader // underlying reader
	N int64     // max bytes remaining
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func main() {
	buf := bytes.NewReader([]byte("test message"))
	rd := LimitedReader{
		R: buf,
		N: 4,
	}
	bts := make([]byte, 10)
	n, err := rd.Read(bts)
	fmt.Println(err)
	fmt.Printf("%s\n", bts[:n])
	fmt.Println(n)
	fmt.Println(len(bts))
}
