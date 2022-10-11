package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Data struct {
	Foo TrimSpaceString `json:"foo" binding:"required"`
}

type TrimSpaceString string

func (t *TrimSpaceString) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, "\"")
	data = bytes.Trim(data, " ")
	*t = TrimSpaceString(strings.TrimSpace(string(data)))
	return nil
}

func main() {
	str := `{"foo": "   b ar   "}`
	var d Data
	json.Unmarshal([]byte(str), &d)
	fmt.Println(d)
}
