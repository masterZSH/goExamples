package main

import (
	"fmt"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

func TestLowerCaseReader(t *testing.T) {
	r := LowerCaseReader{reader: strings.NewReader("aAaAaAbBbBbB")}
	p := make([]byte, 1024)
	n, err := r.Read(p)
	fmt.Println(p[:n], n, err)
	assert.Equal(t, p[:n], []byte("aaabbb"))

	err = iotest.TestReader(
		&LowerCaseReader{reader: strings.NewReader("aBcDeFgHiJ")},
		[]byte("acegi"),
	)
	if err != nil {
		t.Fatal(err)
	}

}
