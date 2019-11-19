package main

import "fmt"

// 定义一个结构体类型 Base，它包含一个字段 id，
// 方法 Id() 返回 id，方法 SetId() 修改 id。
// 结构体类型 Person 包含 Base，
// 及 FirstName 和 LastName 字段。
// 结构体类型 Employee 包含一个 Person
//  和 salary 字段。
//创建一个 employee 实例，然后显示它的 id。

// Base include id
type Base struct {
	id int
}

// Person declare Person struct
type Person struct {
	FirstName, LastName string
	b                   *Base
}

type Employee struct {
	p      *Person
	salary float64
}

// SetID set id
func (b *Base) SetID(i int) {
	b.id = i
}

func main() {
	var b *Base = &Base{1}
	var p *Person = &Person{"foo", "bar", b}

	var e *Employee = new(Employee)
	e.p = p
	e.salary = 1.1
	e.p.b.SetID(33)
	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", e.p.b)
}
