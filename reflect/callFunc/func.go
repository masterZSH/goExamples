package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name, Addr string
}

func (u *User) GetName() (string, error) {
	fmt.Printf("%s,\n", u.Name)
	return u.Name, nil
}

func NewUser() *User {
	return new(User)
}

func main() {
	u := NewUser()
	u.Name = "12345"
	getValue := reflect.ValueOf(u)
	methodName := getValue.MethodByName("GetName")
	args := []reflect.Value{}
	name := methodName.Call(args)
	fmt.Println(name)
}
