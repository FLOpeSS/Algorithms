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

func recPrint(node *nodeTree) {
	if node != nil {
		fmt.Printf("%d", node.Data)
	}

	recPrint(node.Right)
	recPrint(node.Left)
}

func printNodes(root *nodeTree) {
	if root == nil {
		return
	}

	stack := []*nodeTree{}

	current := root

	for current != nil || len(stack) > 0 {
		stack = append(stack, current)
		current = current.Left
	}

	current = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
}

func main() {
	node := newNode(5)
	node.Left = newNode(6)
	node.Right = newNode(7)
	recPrint(node)
}
