package main

import (
	"strings"
	"testing"
	"testing/iotest"
)

func TestLowerCaseReader(t *testing.T) {
	// r := &LowerCaseReader{
	// 	reader: strings.NewReader("aBcDeFgHiJ"),
	// 	data:   make([]byte, 1024),
	// }
	// data, err := io.ReadAll(r)
	// fmt.Println(data, err)
	// p := make([]byte, 3)
	// n, err := r.Read(p)
	// fmt.Println(n, err, p)

	// const msg = "Now is the time for all good gophers."

	// rs := strings.NewReader(msg)
	// if err := iotest.TestReader(rs, []byte(msg)); err != nil {
	// 	t.Fatal(err)
	// }

	err := iotest.TestReader(
		&LowerCaseReader{
			reader: strings.NewReader("aBcDeFgHiJ"),
			data:   make([]byte, 1024),
		},
		[]byte("acegi"),
	)
	if err != nil {
		t.Fatal(err)
	}
}
