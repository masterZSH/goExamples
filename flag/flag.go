package main

import (
	"flag"
)

var (
	h, H, help bool
	v, V       bool
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&H, "H", false, "this help")
	flag.BoolVar(&help, "help", false, "this help")

	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")
}

func main() {
	flag.Parse()

	if h {
		flag.PrintDefaults()
	}
}
