package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	fmt.Print(err)
	fmt.Printf("%v", f)
}
