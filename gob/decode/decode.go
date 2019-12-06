package main

import (
	"bufio"
	"encoding/gob"
	"log"
	"os"
)

// FileName decode file
const FileName string = "vcard.gob"

// 写一个程序读取 vcard.gob 文件，解码并打印它的内容。

func main() {
	// f, err := os.Open(FileName)
	f, err := os.Open(FileName)
	defer f.Close()
	if err != nil {
		log.Fatal("ERROR", err)
	}
	inReader := bufio.NewReader(f)
	dc := gob.NewDecoder(inReader)
	var v interface{}
	dc.Decode(&v)
	log.Printf("%v", dc)
}
