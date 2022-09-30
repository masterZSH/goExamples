package main

import (
	"io"
	"unicode/utf8"
)

type LowerCaseReader struct {
	reader io.Reader
}

func (l LowerCaseReader) Read(p []byte) (int, error) {

	n, err := l.reader.Read(p)
	if err != nil {
		return 0, err
	}
	r := make([]byte, 0)
	for _, v := range p[:n] {
		if isLower(v) {
			r = append(r, v)
		}
	}
	copy(p, r)
	return len(r), nil
}

func isLower(b byte) bool {
	isASCII := true
	if b >= utf8.RuneSelf {
		isASCII = false
	}
	if !isASCII {
		return false
	}

	if 'a' <= b && b <= 'z' {
		return true
	}
	return false
}
