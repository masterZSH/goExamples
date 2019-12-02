package main

import (
	"fmt"
	"io/ioutil"
)

// Page struct
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return ioutil.WriteFile(p.Title, p.Body, 0644)
}

func load(p *Page) ([]byte, error) {
	return ioutil.ReadFile(p.Title)
}

func main() {
	p := &Page{"copy", []byte{'1', '2'}}
	p.save()
	p.Title = "test"
	content, _ := load(p)
	fmt.Printf("%s", content)
}
