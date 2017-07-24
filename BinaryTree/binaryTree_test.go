package binaryTree

import (
	"reflect"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	if reflect.TypeOf(*NewBinaryTree()).Name() != `BinaryTree` {
		t.Error(`Constructor BinaryTree error`)
	}
}
func TestBinaryTree_Add(t *testing.T) {
	tree := NewBinaryTree()
	tree.Add(2)
	if tree.Count() != 1 {
		t.Error(`tree method Add not work`)
	}
}

func TestBinaryTree_Remove(t *testing.T) {
	tree := NewBinaryTree()
	tree.Add(2)
	tree.Add(1)
	tree.Add(4)
	tree.Remove(4)
	if tree.Count() != 2 {
		t.Error(`tree method remove not work`)
	}
	if tree.head.right != nil {
		t.Error(`tree method remove not work`)
	}
}

func TestBinaryTree_Contains(t *testing.T) {
	tree := NewBinaryTree()
	tree.Add(2)
	tree.Add(1)
	tree.Add(4)

	if tree.Contains(4) != true {
		t.Error(`tree method contains not work`)
	}
}

func TestBinaryTree_Count(t *testing.T) {
	tree := NewBinaryTree()
	tree.Add(2)
	tree.Add(1)
	tree.Add(4)
	if reflect.TypeOf(tree.Count()).Kind() != reflect.Int {
		t.Error(`tree method Count not work`)
	}
}

func TestBinaryTree_Clear(t *testing.T) {
	tree := NewBinaryTree()
	tree.Add(2)
	tree.Add(3)

	if tree.Clear(); tree.Count() != 0 {
		t.Error(`tree method Clear not work`)
	}
}
