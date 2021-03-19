package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

const (
	sectionSize = 4 << 10
)

func main() {
	hashFile("test.file")
	// hashFile("large.file")
}

func hashFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fs, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	sectionReader := io.NewSectionReader(f, 0, fs.Size())
	hash := sha256.New()
	sz := sectionReader.Size()
	if sz < sectionSize {
		buff := make([]byte, sz)
		sectionReader.Read(buff)
		hash.Write(buff)
	} else {
		buff := make([]byte, sectionSize)
		sectionReader.Read(buff)
		hash.Write(buff)
		carry := sz / sectionSize
		rand.Seed(carry)
		rd := rand.Int63n(sz - sectionSize)
		sectionReader.Seek(rd, 1)
		sectionReader.Read(buff)
		hash.Write(buff)
	}
	hash.Write([]byte(fmt.Sprint(sectionReader.Size())))
	sum := hash.Sum(nil)
	fmt.Printf("%x\n", sum)
}
