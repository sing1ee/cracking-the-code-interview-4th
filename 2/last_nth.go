package main

import (
	"fmt"
)

type LinkedListNode struct {
	next *LinkedListNode
	value int
}

func FindNthLast(head *LinkedListNode, n int) *LinkedListNode {
	pFirst := head
	pLast := head

	i := 0
	for ; i < n; i++ {
		if nil == pFirst {
			break
		}
		pFirst = pFirst.next
	}

	if i != n {
		return nil
	}

	for ;nil != pFirst; {
		pFirst = pFirst.next
		pLast = pLast.next
	}
	return pLast
}

func Case1() {
	head := new(LinkedListNode)
	head.value = 0
	last := head
	for i := 1; i < 10; i++ {
		node := new(LinkedListNode)
		node.value = i
		last.next = node
		last = node
	}
	fmt.Println(FindNthLast(head, 5).value)
	fmt.Println(FindNthLast(head, 10).value)
	fmt.Println(FindNthLast(head, 11).value)
}

func main() {
	Case1()
}
