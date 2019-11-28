package main

import "fmt"

var (
	firstName, lastName string
)

func main() {
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels

}
