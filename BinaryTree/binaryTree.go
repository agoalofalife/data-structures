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

// helper recursive method for add value in tree
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

	current, parent := tree.findWithParent(value)

	if current == nil {
		return false
	}
	tree.count--

	if current.left == nil && current.right == nil {
		if parent.left.value == current.value {
			parent.left = nil
		} else if parent.right.value == current.value {
			parent.right = nil
		}
	} else if current.right == nil || current.left == nil {
		current = current.right
	} else if current.right != nil && current.left != nil {

		leftmost := current.right
		for leftmost.left != nil {
			leftmost = leftmost.left
		}

		if parent.left.value == current.value {
			parent.left = leftmost
		} else if parent.right.value == current.value {
			parent.right = leftmost
		}
		parent.left = leftmost
	}

	return true
}

func (tree *BinaryTree) Contains(value int) bool {
	finder, _ := tree.findWithParent(value)
	return finder != nil
}
func (tree *BinaryTree) Count(value int) int {
	return tree.count
}

func (tree *BinaryTree) Clear(value int) {
	tree.head = nil
	tree.count = nil
}

// find value position and return value node and his parent node
func (tree *BinaryTree) findWithParent(value int) (current *Node, parent *Node) {
	current = tree.head
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
	return
}

func main() {
	tree := NewBinaryTree()
	tree.Add(8)
	tree.Add(5)
	tree.Add(10)
	tree.Add(2)
	tree.Add(7)
	tree.Add(6)
	tree.Remove(5)
	fmt.Println(tree.head.left)
	os.Exit(0)
}
