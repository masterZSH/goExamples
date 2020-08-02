package init

import "fmt"

var Mod string

func init() {
	Mod = "12"
}

func PrintLog() {
	fmt.Print(Mod)
}
