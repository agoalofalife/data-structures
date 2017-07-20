package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := Stack{}
	stack.Push(2)

	if stack.Peek() != 2 {
		t.Error("Stack not have value")
	}
}

func TestPop(t *testing.T) {
	stack := Stack{}
	stack.Push(2)

	if stack.Pop() != 2 {
		t.Error("Stack not have value")
	}

	if stack.Count() != 0 {
		t.Error("Stack not right count")
	}
}

func TestPopExpectedPanic(t *testing.T)  {
	defer func() {
		str := recover()
		if str != "Stack is empty" {
			t.Fatalf("Wrong panic message: %s", str)
		}
	}()
	stack := Stack{}
	stack.Pop()
}

func TestCount(t *testing.T) {
	stack := Stack{}
	stack.Push(2)

	if stack.Count() != 1 {
		t.Error("Stack not right count")
	}

}

func TestPeekExpectedPanic(t *testing.T) {
	defer func() {
		str := recover()
		if str != "Stack is empty" {
			t.Fatalf("Wrong panic message: %s", str)
		}
	}()
	stack := Stack{}
	stack.Peek()
}

func TestPeek(t *testing.T) {
	stack := Stack{}
	stack.Push(2)
	if stack.Peek() != 2 {
		t.Error("Wrong get result from Peek")
	}
}