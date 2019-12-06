package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Name, Addr string
}

type VCard struct {
	FirstName int
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	a := &Address{"1", "2"}
	b := &Address{"3", "4"}
	p := VCard{123, "on", []*Address{a, b}, "abc"}
	fmt.Printf("%v", p)
	js, _ := json.Marshal(p)
	fmt.Printf("json :%s", js)
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(p)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
