package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("foo r1")
	r2 := strings.NewReader("foo r2")
	r3 := strings.NewReader("foo r3")

	r := io.MultiReader(r1, r2, r3)
	io.Copy(os.Stdout, r)


}
