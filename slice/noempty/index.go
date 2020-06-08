package main

import "fmt"

func main() {
	data := []string{"test", "", "co"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", noempty(data))

}

func noempty(data []string) []string {
	temp := data[:0]
	for _, item := range data {
		if item != "" {
			temp = append(temp, item)
		}
	}
	return temp
}
