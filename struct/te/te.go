package main

import "fmt"

type Pooler interface{

}

type T struct{

}

func main() {
	var _ Pooler = (*T)(nil)
	fmt.Print(Pooler)
}