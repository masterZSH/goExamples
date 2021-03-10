package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strings.NewReader("1233")
	r := bufio.NewReader(s)
	// r.Read([]byte{'1'})
	fmt.Println(r.Buffered(), r.Size())

	rs := bufio.NewReaderSize(s, 100)
	rs.WriteTo(os.Stdout)

	// 123
	st, _ := r.Peek(3)
	fmt.Printf("%s", st)

	// scanner
	rd := strings.NewReader("test commm")
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	write()
}

func write() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	w.Write([]byte{'a'})
	w.WriteByte('b')
	w.WriteRune('c')
	w.WriteString("d")
	fmt.Fprintf(w, "%s", "test1")
	fmt.Fprintf(w, "%s", "test2")

}
