package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("n", false, "print new line")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	// who := "alice"
	// if len(os.Args) > 1 {
	// 	who += strings.Join(os.Args[1:], " ")
	// }
	// fmt.Println(who)

	// helloStr := "Hello "
	// if len(os.Args) > 1 {
	// 	helloStr += strings.Join(os.Args[1:], " ")
	// 	// add end tag
	// 	helloStr += "!"
	// }
	// fmt.Println(helloStr)

	flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
