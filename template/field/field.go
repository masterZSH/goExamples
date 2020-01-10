package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name                string
	nonExportedAgeField string
}

func main() {
	t := template.New("hello")
	t, _ = t.Parse("hello {{.Name}}!")
	p := Person{Name: "test", nonExportedAgeField: "las"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Print("error:%v", err)
	}
}
