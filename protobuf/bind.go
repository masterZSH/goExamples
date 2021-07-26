package main

import (
	"fmt"
	"main/testdata"

	"github.com/golang/protobuf/proto"
)

func main() {

	b := []byte("")
	test := &testdata.Test{
		Label: proto.String("foo"),
	}
	data, _ := proto.Marshal(test)
	fmt.Printf("%s\n", data)

	err := proto.Unmarshal(b, test)
	fmt.Println(err, b, test)
}
