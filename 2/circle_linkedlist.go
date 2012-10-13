package main

import (
	"fmt"
)

type LinkedListNode struct {
	next *LinkedListNode
	value int
}

func Find(head *LinkedListNode) *LinkedListNode {
	if nil == head {
		return nil
	}
	oneStep := head
	twoStep := head
	for ; ; {
		oneStep = oneStep.next
		twoStep = twoStep.next.next
		if oneStep.value == twoStep.value {
			break
		}
	}
	oneStep = head
	for ; oneStep.value != twoStep.value; {
		oneStep = oneStep.next
		twoStep = twoStep.next
	}
	return oneStep
}

func main() {
	// arr := []int{1, 2, 3, 4, 5, 3}
	head := new(LinkedListNode)
	head.value = 1

	node1 := new(LinkedListNode)
	node1.value = 2
	head.next = node1

	node2 := new(LinkedListNode)
	node2.value = 3
	node1.next = node2
	
	node3 := new(LinkedListNode)
	node3.value = 4
	node2.next = node3

	node4 := new(LinkedListNode)
	node4.value = 5
	node3.next = node4

	node4.next = head

	fmt.Println(Find(head).value)
}