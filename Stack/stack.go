//---------------------------------------------------------------------------
//-[ Stack ]-----------------------------------------------------------------
//---------------------------------------------------------------------------
//---------------------------------------------------------------------------

package main

import (
	"os"
	"log"
)

type Stack struct {
	stack []int
	count int
}
// push value stack
func (stack *Stack) Push(value int){
	stack.stack = append(stack.stack, value)
	stack.count++
}

// pop value from stack
func (stack *Stack) Pop() (value int) {
	if stack.count == 0 {
		panic("Stack is empty")
	}
	target := len(stack.stack) - 1
	value = stack.stack[target:][0]
	stack.stack = stack.stack[:target]
	stack.count--
	return
}

func (stack *Stack) Count() int {
	return stack.stack[:1][0]
}

func main()  {
	s := Stack{}
	s.Push(2)
	s.Push(1)
	s.Push(3)
	s.Pop()
	log.Println(s)
	os.Exit(2)
}