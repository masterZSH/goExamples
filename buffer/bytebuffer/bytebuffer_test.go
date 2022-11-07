package bytebuffer

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	var b ByteBuffer
	n, err := b.Write([]byte("123"))
	fmt.Println(n, err)

	rd := bytes.NewReader([]byte("456"))
	rn, err := b.ReadFrom(rd)
	fmt.Println(rn, err)

}
