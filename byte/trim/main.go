package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/leebenson/conform"
)

type Data struct {
	Foo TrimSpaceString `json:"foo" binding:"required"`
}

type Data2 struct {
	Foo string `json:"foo" binding:"required" conform:"trim"`
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

	d2 := Data2{
		Foo: "   b ar   ",
	}
	conform.Strings(&d2)
	fmt.Println(d2)
}
