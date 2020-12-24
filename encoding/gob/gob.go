package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Kv struct {
	Key string
	Val string
}

type Foo struct {
}

func main() {
	var r bytes.Buffer
	kv := Kv{
		Key: "foo",
		Val: "bar",
	}
	err := gob.NewEncoder(&r).Encode(kv)
	if err != nil {
		log.Fatal(err)
	}

	// to string
	data := r.String()

	var dkv Kv
	// decode
	gob.NewDecoder(bytes.NewBufferString(data)).Decode(&dkv)

	fmt.Print(kv, dkv)
}
