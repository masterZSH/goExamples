package main

import (
	"fmt"
	"strings"
	"time"
)

type identifier struct {
	name int
	v    int
	f    func(int) int
}

func main() {
	var i identifier
	i.v = 3
	fmt.Printf("%+v\n", i)

	var mi *identifier = new(identifier)
	mi.v = 1
	fmt.Printf("%+v\n", mi)

	ms := &identifier{}
	fmt.Printf("%+v\n", ms)

	var p Person
	p.firstName = "foo"
	p.lastName = "bar"
	Up(&p)
	fmt.Printf("%+v\n", p)

	var np = new(Person)
	np.firstName = "f"
	np.lastName = "b"
	Up(np)
	fmt.Printf("%+v\n", np)

	var np1 = &Person{"fo", "la"}
	Up(np1)
	fmt.Printf("%+v\n", np1)
}

// Node a Node struct
type Node struct {
	data float64
	su   *Node
}

// Person struct
type Person struct {
	firstName string
	lastName  string
}

// Up change Person name to Up
func Up(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

type Vcard struct {
	name     string
	birth    time.Time
	addreses map[string]*Address
}

type Address struct {
	addr string
}
