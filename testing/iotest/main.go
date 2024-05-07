package main

import (
	"io"
)

type LowerCaseReader struct {
	reader io.Reader
	data   []byte
	i      int
}

func (l *LowerCaseReader) Read(p []byte) (n int, err error) {
	if l.i >= len(l.data) {
		return 0, io.EOF
	}
	if len(p) == 0 {
		return 0, nil
	}
	var nn int
	for {
		nn, err = l.reader.Read(l.data)
		if err != nil {
			return len(p), err
		}
		var temp []byte
		for _, c := range l.data[0:nn] {
			switch {
			case c >= 'a' && c <= 'z':
				temp = append(temp, c)
				n++
				continue
			default:
				continue
			}
		}
		p = p[0:len(temp)]
		copy(p, temp)
		l.i += len(p)
		l.data = p
	}
	return
}
