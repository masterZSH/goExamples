package main

import (
	"fmt"

	"google.golang.org/grpc/test/codec_perf"

	"github.com/golang/protobuf/proto"
)

func Marshal(v interface{}) ([]byte, error) {
	vv, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("failed to marshal, message is %T, want proto.Message", v)
	}
	return proto.Marshal(vv)
}

func Unmarshal(data []byte, v interface{}) error {
	vv, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("failed to unmarshal, message is %T, want proto.Message", v)
	}
	return proto.Unmarshal(data, vv)
}

func main() {
	p := &codec_perf.Buffer{}
	p.Body = []byte("123")
	r, err := Marshal(p)
	fmt.Println(r, err)

	var d codec_perf.Buffer
	err = Unmarshal(r, &d)
	fmt.Println(d, err)
}
