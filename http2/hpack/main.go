package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/net/http2/hpack"
)

func main() {
	var buf bytes.Buffer
	enc := hpack.NewEncoder(&buf)
	enc.WriteField(hpack.HeaderField{
		Name:  "abc",
		Value: "123",
	})
	enc.SetMaxDynamicTableSize(255)

	d := hpack.NewDecoder(4096, func(_ hpack.HeaderField) {})
	defer d.Close()
	_, err := d.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	arr, err := d.DecodeFull(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arr)
}
