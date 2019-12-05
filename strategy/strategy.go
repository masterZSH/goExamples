package main

import (
	"flag"
	"fmt"
	"os"
)

type Printer interface {
	Print()
}

type A struct {
}

func (a *A) Print() {
	fmt.Fprintf(os.Stdout, "aaaa")
}

type B struct {
}

func (b *B) Print() {
	fmt.Fprintf(os.Stdout, "bbbb")
}

func NewPrint(arg string) Printer {
	if arg == "a" {
		return &A{}
	} else {
		return &B{}
	}
}

func main() {
	flag.Parse()
	var f = flag.NArg()
	if f == 0 {
		return
	}
	p := NewPrint(flag.Arg(0))
	p.Print()
}
