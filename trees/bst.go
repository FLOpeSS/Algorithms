package main

import "fmt"

type nodeTree struct {
	Data  int
	Left  *nodeTree
	Right *nodeTree
}

func newNode(data int) *nodeTree {
	return &nodeTree{
		Data:  data,
		Left:  nil,
		Right: nil,
	}

}

func printNodes(root *nodeTree) {
	stack := []*nodeTree{
		{Data: 1, Left: nil, Right: nil},
		{Data: 2, Left: nil, Right: nil},
	}

	random := 40
	pointer := &random
	fmt.Println(*pointer)
}

func main() {
	node := newNode(5)
	printNodes(node)
}
