package main 

import (
	"fmt"
)

type LinkedListNode struct {
	next *LinkedListNode
	value int
}

func SumLinkedList(left, right, re *LinkedListNode, carry int) {
	if nil == left && nil == right {
		if 1 == carry {
			node := new(LinkedListNode)
			node.value = carry
			re.next = node
		}
		return
	}
	if nil == left {
		if 1 == carry {
			carryNode := new(LinkedListNode)
			carryNode.value = 1
			SumLinkedList(carryNode, right, re, 0)
		} else {
			re.next = right
		}
		return
	}
	if nil == right {
		if 1 == carry {
			carryNode := new(LinkedListNode)
			carryNode.value = 1
			SumLinkedList(left, carryNode, re, 0)
		} else {
			re.next = left.next
		}
		return
	}
	sum := left.value + right.value + carry
	newCarry := 0
	if sum > 9 {
		newCarry = 1
		sum = sum - 10
	}
	node := new(LinkedListNode)
	node.value = sum
	re.next = node
	SumLinkedList(left.next, right.next, node, newCarry)
}

func BuildLinedList(n int) *LinkedListNode {
	if n < 0 {
		return nil
	}
	head := new(LinkedListNode)
	head.value = 0
	last := head
	for i := 1; i < n; i++ {
		node := new(LinkedListNode)
		node.value = i
		last.next = node
		last = node
	}
	return head
}

func main() {
	left := BuildLinedList(9)
	right := BuildLinedList(10)
	sum := new(LinkedListNode)
	sum.value = -1
	SumLinkedList(left, right, sum, 0)
	for e := sum.next; e != nil; e = e.next {
		fmt.Print(e.value, "->")
	}
	fmt.Println("nil")
}