package main

import (
	"fmt"
)

// Stest struct
type Stest struct {
	f float64
	string
	int
}

func main() {
	var sT = &Stest{1.1, "2", 3}
	fmt.Printf("%+v\n", sT)
}
