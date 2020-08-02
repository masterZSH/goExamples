package main

import (
	"fmt"
	i "init"
)

func init() {
	fmt.Printf("%s\n", i.Mod)
	i.Mod = "123"

}

func main() {
	i.PrintLog()
}
