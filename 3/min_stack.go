package main

import (
	"fmt"
	"container/list"
)

/*
*How would you design a stack which, in addition to push and pop, also has a function 
*min which returns the minimum element? Push, pop and min should all operate in  O(1) time 
*/

type MinStack struct {
	value *list.List
	minValues *list.List
}


func (stack *MinStack) Push(value int) {
	stack.value.PushBack(value)
	if stack.minValues.Len() == 0 {
		stack.minValues.PushBack(value)
	} else {
		if value < stack.minValues.Back().Value.(int) {
			stack.minValues.PushBack(value)
		}
	}
}

func (stack *MinStack) Pop() int {
	value := stack.value.Back();
	ret := value.Value.(int)
	stack.value.Remove(value)
	if ret == stack.minValues.Back().Value.(int) {
		stack.minValues.Remove(stack.minValues.Back())
	}
	return ret
}

func (stack *MinStack) Peek() int {
	return stack.value.Back().Value.(int)
}

func (stack *MinStack) Min() int {
	return stack.minValues.Back().Value.(int)
}

func (stack *MinStack) IsEmpty() bool {
	return 0 == stack.value.Len()
}

func main() {
	stack := &MinStack{list.New(), list.New()}
	for i := 100; i >= 0; i-- {
		stack.Push(i)
	}
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
}



