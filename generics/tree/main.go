package main

type Node[v any] struct {
	Data  v
	Left  *Node[v]
	Mid   *Node[v]
	Right *Node[v]
}

type Tree[v any] struct {
	Root *Node[v]
}

func NewTree[v any]() *Tree[v] {
	return &Tree[v]{}
}

func main() {
	tree := NewTree[int]()

	root := &Node[int]{
		Data: 1,
	}
	root.Left = &Node[int]{
		Data: 2,
	}
	tree.Root = root
}
