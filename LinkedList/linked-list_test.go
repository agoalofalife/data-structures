package main

import (
	"testing"
	"reflect"
)

func TestNewNode(t *testing.T)  {
	node := NewNode(2)

	if reflect.TypeOf(*node).Name() != `Node`{
		t.Error(`Constructor not return Node struct`)
	}
}

func TestAdd(t *testing.T)  {
	list := LinkedList{}
	list.Add(2)

	if list.Count() != 1  {
		t.Error(`Add method not work`)
	}
	if list.IndexOf(2) != 0 {
		t.Error(`Wrong method not work for indexOf`)
	}
}

func TestRemove(t *testing.T)  {
	list := LinkedList{}
	list.Add(2)
	list.Add(1)
	list.Add(3)

	list.Remove(3)

	if list.Count() != 2{
		t.Error(`Wrong count after call remove method`)
	}

	if list.Exist(3) != false {
		t.Error(`Wrong exist after call remove method`)
	}
}

func TestExist(t *testing.T)  {
	list := LinkedList{}
	list.Add(2)

	if list.Exist(2) != true {
		t.Error(`Method exist not work`)
	}
}

func TestReset(t *testing.T)  {
	list := LinkedList{}
	list.Add(2)
	list.Add(2)
	list.Add(2)
	list.Reset()

	if list.Count() != 0 {
		t.Error(`Method reset not work`)
	}
}

func TestCopyTo(t *testing.T)  {
	list := LinkedList{}
	list.Add(1)
	list.Add(2)

	tempArr := make([]int, 2)
	list.CopyTo(tempArr, 0)

	if len(tempArr) != 2 {
		t.Error(`Method CopyTo not work`)
	}
}

func TestCount(t *testing.T)  {
	list := LinkedList{}
	list.Add(1)
	list.Add(2)

	if list.Count() != 2 {
		t.Error(`Method Count not right`)
	}
}

func TestIndexOf(t *testing.T)  {
	list := LinkedList{}
	list.Add(1)
	list.Add(2)

	if list.IndexOf(2) != 1{
		t.Error(`Method indexOf not true index`)
	}
}

func TestGetByIndex(t *testing.T)  {
	list := LinkedList{}
	list.Add(1)
	list.Add(2)

	if list.GetByIndex(1) != 2{
		t.Error(`Position index not true`)
	}
}

func ExamplePrint() {
		nodeFirst := &Node{1, nil}
		nodeSecond := &Node{2, nil}
		nodeFirst.next = nodeSecond
		Print(nodeFirst)
	// Output: 1
}
// go test -v . | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''