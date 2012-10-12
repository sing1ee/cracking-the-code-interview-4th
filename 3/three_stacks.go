package main

import (
	"fmt"
)

type StackNode struct {
	previous int
	value int
}

const stackSize = 300

type ThreeStack struct {
	indexUsed int
	stackPointer []int
	buffer []*StackNode
}

func (stack *ThreeStack) PrintStack(size int) {
	for i := 0; i < size; i++ {
		if nil != stack.buffer[i] {
			fmt.Print(stack.buffer[i].value, "-")
		} else {
			fmt.Print("nil", "-")
		}
	}
	fmt.Println(stack.indexUsed)
}

func (stack *ThreeStack) Push(stackNum, value int) {
	lastIndex := stack.stackPointer[stackNum]
	stack.stackPointer[stackNum] = stack.indexUsed
	stack.indexUsed++
	stack.buffer[stack.stackPointer[stackNum]] = &StackNode{lastIndex, value}
}

func (stack *ThreeStack) Pop(stackNum int) int {
	lastIndex := stack.stackPointer[stackNum]
	value := stack.buffer[lastIndex].value
	stack.stackPointer[stackNum] = stack.buffer[lastIndex].previous
	stack.buffer[lastIndex] = nil
	// stack.indexUsed--
	return value
}

func (stack *ThreeStack) Peek(stackNum int) int {
	return stack.buffer[stack.stackPointer[stackNum]].value
}

func (stack *ThreeStack) IsEmpty(stackNum int) bool {
	return -1 == stack.stackPointer[stackNum]
}

func Case1() {
	stack := &ThreeStack{0, []int{-1, -1, -1}, make([]*StackNode, stackSize * 3)}
	stackNum := 0
	fmt.Println("push")
	for i := 0; i < 10; i++ {
		sn := stackNum % 3
		stack.Push(sn, i)
		fmt.Println(sn, i)
		stackNum++
	}
	stackNum = 0
	fmt.Println("pop")
	for i := 0; i < 10; i++ {
		sn := stackNum % 3
		fmt.Println(sn, stack.Pop(sn))
		stackNum++
	}
}

func Case2() {
	stack := &ThreeStack{0, []int{-1, -1, -1}, make([]*StackNode, stackSize * 3)}
	stack.Push(0, 0)
	stack.Push(0, 1)
	stack.Push(0, 2)
	stack.Push(1, 0)
	stack.Push(1, 1)
	stack.Push(1, 2)
	stack.PrintStack(10)
	fmt.Println(0, stack.Pop(0), 2)
	fmt.Println(1, stack.Pop(1), 2)
	stack.PrintStack(10)
	stack.Push(2, 0)
	stack.PrintStack(10)
	// stack.Push(2, 1)
	// stack.Push(2, 2)
	// fmt.Println(2, stack.Pop(2), 2)
	// fmt.Println(0, stack.Pop(0), 1)
	// stack.Push(1, 3)
	fmt.Println(1, stack.Pop(1), 1)
}

func main() {
	Case2()
}