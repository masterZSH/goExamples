package main
import (
	"fmt"
)

const (
	i1 = iota
	i2 
	i3 
	i4 
)


func main() {
	ch := make(chan int)
	go run(ch)
	for i := range ch {
		fmt.Printf("%d ", i)
	}
}


func run(ch chan int){
	for i:=0;i<1000;i++{
		ch<-i;
	}
	close(ch)
}