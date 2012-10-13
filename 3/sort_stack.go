package main

import (
	"fmt"
	"container/list"
)
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
	return topEl.Value
}

func (stack *Stack) Peek() interface{} {
	return stack.elements.Back().Value
}

func (stack *Stack) IsEmpty() bool {
	return 0 == stack.Size()
}

func (stack *Stack) Size() int {
	return stack.elements.Len()
}

func SortStack(stack *Stack, depth int) {
	if depth < 1 {
		return
	}
	tmp := depth
	tmpStack := &Stack{list.New()}
	max := -1
	for ; tmp > 0; tmp-- {
		value := stack.Pop().(int)
		if value > max {
			max = value
		}
		tmpStack.Push(value)
	}
	stack.Push(max)
	tmp = depth
	first := false // for the duplicate element
	for ; tmp > 0; tmp-- {
		value := tmpStack.Pop().(int)
		if value == max && !first{
			first = true
		} else {
			stack.Push(value)
		}
	}
	SortStack(stack, depth - 1)
}

func main() {
	stack := &Stack{list.New()}
	for i := 100; i > 0; i-- {
		stack.Push(i)
	}
	SortStack(stack, 100)
	for ; !stack.IsEmpty(); {
		fmt.Println(stack.Pop().(int))
	}
}

