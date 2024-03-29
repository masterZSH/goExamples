package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	errInvalidWrite = errors.New("invalid write")
)

func copyBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	// if wt, ok := src.(io.WriterTo); ok {
	// 	return wt.WriteTo(dst)
	// }
	// // Similarly, if the writer has a ReadFrom method, use it to do the copy.
	// if rt, ok := dst.(io.ReaderFrom); ok {
	// 	return rt.ReadFrom(src)
	// }
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = errInvalidWrite
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}

func main() {
	var dst bytes.Buffer
	// var src bytes.Buffer
	src := bytes.NewReader([]byte("123"))

	time.Sleep(100 * time.Millisecond)
	copyBuffer(&dst, src, nil)

	fmt.Printf("%s\n", dst.String())

	//
	reader, writer := io.Pipe()

	go func() {
		// fmt.Fprint(writer, "some io.Reader stream to be read\n")
		// writer.Close()
		writer.Write([]byte("1234"))
		writer.Close()
	}()

	if _, err := io.Copy(os.Stdout, reader); err != nil {
		log.Fatal(err)
	}

}
