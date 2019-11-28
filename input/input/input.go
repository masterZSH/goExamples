package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	firstName, lastName string
)

func main() {
	// fmt.Scanln(&firstName, &lastName)
	// fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	// ReadInput()
	WordLetterCount()
}

func ReadInput() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("input :")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Print(input)
	}
}

var (
	chars, fields, lines int
)

func WordLetterCount() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input ")
	chars, fields, lines = 0, 0, 0
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(0)
		}
		if input == "S\r\n" { // Windows, on Linux it is "S\n"
			fmt.Println("Here are the counts:", chars, fields, lines)
			os.Exit(0)
		}
		Count(input)
	}
}

func Count(input string) {
	chars += len(input) - 2
	fields += len(strings.Fields(input))
	lines++
}
