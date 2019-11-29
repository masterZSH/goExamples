package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	name  string
	price float64
	stock int
}

func main() {
	file, _ := os.Open("test.text")
	csvReader := bufio.NewReader(file)
	var books []Book = make([]Book, 1)
	defer file.Close()
	for {
		line, err := csvReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = string(line[:len(line)-2])
		if line == "" {
			continue
		}
		var book *Book = new(Book)
		var fields = strings.Split(line, ";")
		book.name = fields[0]
		price, err := strconv.ParseFloat(fields[1], 64)
		book.price = price
		book.stock, _ = strconv.Atoi(fields[2])
		if books[0].name == "" {
			books[0] = *book
		} else {
			books = append(books, *book)
		}
	}
	fmt.Print(books)
}
