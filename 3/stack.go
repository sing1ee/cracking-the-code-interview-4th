package stack

import (
	"container/list"
)
/*
*implement stack by double linked list
*/

type Stack struct {
	elements *list.List
}

func (stack *Stack) Elements() *list.List {
	return stack.elements
}

func (stack *Stack) Push(value interface{}) {
	stack.elements.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	topEl := stack.elements.Back()
	stack.elements.Remove(topEl)
	return topEl
}

func (stack *Stack) Peek() interface{} {
	return stack.elements.Back()
}

func (stack *Stack) IsEmpty() bool {
	return 0 == stack.elements.Len()
}

func (stack *Stack) Size() int {
	return stack.elements.Len()
}