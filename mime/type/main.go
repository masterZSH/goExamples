package main

import (
	"fmt"
	"mime"
)

func main() {
	fmt.Println(mime.TypeByExtension(".jpg"))
	fmt.Println(mime.TypeByExtension(".csv"))
	fmt.Println(mime.TypeByExtension(".xlsx"))

}
