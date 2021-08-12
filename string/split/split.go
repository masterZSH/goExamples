package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a,b,c,d,e,f,g,h,i,j"

	r := strings.SplitN(str, ",", 1)
	fmt.Println(r)
}
