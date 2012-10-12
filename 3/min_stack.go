package main

import (
	"fmt"
	"math/rand"
	"container/list"
)

/*
*How would you design a stack which, in addition to push and pop, also has a function 
*min which returns the minimum element? Push, pop and min should all operate in  O(1) time 
*/

type MinStack struct {
	top int
	value []int
	minValues *list.List
}


func (stack *MinStack) Push(value int) {
	if stack.top >= len(stack.value)
	stack.value[stack.top] = value
	stack.top++
	if stack.min == 0 {
		stack.minValues.PushBack(value)
	} else {
		if value < stack.minValues.Back().Value.(int) {
			stack.minValues.PushBack(value)
		}
	}
}

func (stack *MinStack) Pop() int {
	value := stack.value[stack.top - 1]
	stack.top--
	if value == stack.minValues.Back().Value.(int) {
		stack.minValues.Remove(stack.minValues.Back())
	}
	return value
}

func (stack *MinStack) Peek() int {
	return stack.value[stack.top - 1]
}

func (stack *MinStack) Min() int {
	return stack.minValues.Back().Value.(int)
}

func (stack *MinStack) IsEmpty() bool {
	return 0 == stack.top
}

func main() {
}