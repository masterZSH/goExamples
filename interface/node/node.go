package main

import (
	"fmt"
)

type Node struct {
	left, right *Node
	data        interface{}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func NewNode(left, right *Node) *Node {
	return &Node{left, right, nil}
}

func main() {
	var root = NewNode(nil, nil)
	root.SetData("root")
	var l = NewNode(nil, root)
	l.SetData("left")
	var r = NewNode(root, nil)
	r.SetData("right")
	fmt.Printf("%+v\n", root)
	fmt.Printf("%+v\n", l)
	fmt.Printf("%+v\n", r)

}
