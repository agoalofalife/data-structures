package main

import (
	"fmt"
	"os"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

// Introduction
// If value Node > then value = 1
// If value Node < then value = -1
// else return 0
func (node *Node) CompareTo(value int) (meter int) {
	if node.value == value {
		return 0
	} else if node.value > value {
		return 1
	} else if node.value < value {
		return -1
	}
	return meter
}

type BinaryTree struct {
	head  *Node
	count int
}

func NewBinaryTree() *BinaryTree {
	tree := &BinaryTree{nil, 0}
	return tree
}

func (tree *BinaryTree) Add(value int) {
	if tree.head == nil {
		tree.head = &Node{value, nil, nil}
	} else {
		tree.addTo(tree.head, value)
	}
	tree.count++
}

func (tree *BinaryTree) addTo(node *Node, value int) {
	if node.value > value {
		if node.left == nil {
			node.left = &Node{value, nil, nil}
		} else {
			tree.addTo(node.left, value)
		}
	} else {
		if node.right == nil {
			node.right = &Node{value, nil, nil}
		} else {
			tree.addTo(node.right, value)
		}
	}
}

func (tree *BinaryTree) Remove(value int) bool {
	parent := Node{}
	current := tree.findWithParent(value, &parent)
	if current == nil {
		return false
	}
	return true
}

func (tree *BinaryTree) findWithParent(value int, parent *Node) *Node {
	current := tree.head
	for current != nil {
		meter := current.CompareTo(value)

		if meter == 1 {
			parent = current
			current = current.left
		} else if meter == -1 {
			parent = current
			current = current.right
		} else {
			break
		}
	}
	//fmt.Println( parent)
	//os.Exit(0)
	return current
}
func main() {
	tree := NewBinaryTree()
	tree.Add(2)
	tree.Add(4)
	tree.Add(1)
	s := &Node{}
	tree.findWithParent(1, s)
	fmt.Println(s)
	os.Exit(0)
}
