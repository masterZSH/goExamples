package main

import (
	"encoding/json"
	"fmt"
)

var j string = "{\"age\":1,\"name\":\"test\",\"items\":[]}"

type test struct {
	age, num int
	items    []string
}

func main() {
	var f interface{}
	var bs []byte = []byte(j)
	json.Unmarshal(bs, &f)
	fmt.Printf("%+v\n", f)

	var t test
	t.items = []string{"123", "456", "789"}
	t.age = 333
	t.num = 222

	js, _ := json.MarshalIndent(t, "", "	")

	fmt.Printf("%+v\n", t)
	fmt.Printf("%s\n", js)
}
