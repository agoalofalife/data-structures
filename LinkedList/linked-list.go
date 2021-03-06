//---------------------------------------------------------------------------
//-[ Linked List ]------------------------------------------------------------
//---------------------------------------------------------------------------
//---------------------------------------------------------------------------

package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

// constructor Node
func NewNode(number int) *Node {
	node := &Node{number, nil}
	return node
}

// Base struct
type LinkedList struct {
	count    int
	head     *Node
	last     *Node
	current  *Node
	previous *Node
}

// Add new node in collection
func (list *LinkedList) Add(value int) {
	node := NewNode(value)

	if list.head == nil {
		list.head = node
		list.last = node
	} else {
		list.last.next = node
		list.last = node
	}
	list.count++
}

// Remove node in collection
func (list *LinkedList) Remove(value int) bool {
	list.previous = nil
	list.current = list.head

	for list.current != nil {
		if list.current.value == value {
			if list.previous != nil {
				list.previous.next = list.current.next
				if list.current.next == nil {
					list.last = nil
				}
			} else {
				list.head = list.head.next
				if list.head == nil {
					list.last = nil
				}
			}
			list.count--
			return true

		}
		list.previous = list.current
		list.current = list.current.next
	}
	return false
}

// Search node in collection
func (list *LinkedList) Exist(value int) bool {
	list.current = list.head
	for list.current != nil {
		if list.current.value == value {
			return true
		}
		list.current = list.current.next
	}
	return false
}

// all clear collection
func (list *LinkedList) Reset() bool {
	list.head = nil
	list.last = nil
	list.count = 0
	return true
}

// Copy values in array with index pass second arguments
func (list *LinkedList) CopyTo(arr []int, value int) {
	list.current = list.head

	for list.current != nil {
		arr[value] = list.current.value
		value++
		list.current = list.current.next
	}
}

// Get count
func (list *LinkedList) Count() int {
	return list.count
}

//get index value
func (list *LinkedList) IndexOf(value int) int {
	list.current = list.head
	index := -1

	for list.current != nil {
		index++
		if list.current.value == value {
			return index
		}
		list.current = list.current.next
	}
	return index
}

// Get value by index
func (list *LinkedList) GetByIndex(index int) int {
	list.current = list.head
	count := 0
	for count < index {
		count++
		list.current = list.current.next
	}
	return list.current.value
}

// print empty linked list, function helper
func Print(node *Node) {
	for node.next != nil {
		fmt.Println(node.value)
		// Output: 1
		node = node.next
	}
}

func main() {
	nodeFirst := &Node{1, nil}
	nodeSecond := &Node{2, nil}
	nodeFirst.next = nodeSecond
	Print(nodeFirst)
}
