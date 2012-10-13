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

type Queue struct {
	inStack, outStack *Stack
}

func (q *Queue) Push(value interface{}) {
	q.inStack.Push(value)
}

func (q *Queue) Pop() interface{} {
	if q.outStack.IsEmpty() {
		for ; !q.inStack.IsEmpty(); {
			q.outStack.Push(q.inStack.Pop())
		}
	}
	return q.outStack.Pop()
}

func (q *Queue) Peek() interface{} {
	if q.outStack.IsEmpty() {
		for ; !q.inStack.IsEmpty(); {
			q.outStack.Push(q.inStack.Pop())
		}
	}
	return q.outStack.Peek()
}

func (q *Queue) IsEmpty() bool {
	return q.inStack.IsEmpty() && q.outStack.IsEmpty()
}

func (q *Queue) Size() int {
	return q.inStack.Size() + q.outStack.Size()
}

func main() {
	q := &Queue{&Stack{list.New()}, &Stack{list.New()}}
	for i := 0; i < 100; i++ {
		q.Push(i)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(i, q.Pop())
	}
}

