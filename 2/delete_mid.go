package main

import (
	"fmt"
)

type LinkedListNode struct {
	next *LinkedListNode
	value int
}

func DelNode(toDelNode *LinkedListNode) {
	if nil == toDelNode || nil == toDelNode.next {
		return
	}

	next := toDelNode.next
	toDelNode.value = next.value
	toDelNode.next = toDelNode.next.next
}

func DeleteMid(head *LinkedListNode) {
	if nil == head {
		return
	}
	pFast := head
	pSlow := head
	pLast := head
	for ; ; {
		pFast = pFast.next
		if nil != pFast {
			pFast = pFast.next
			if nil == pFast {
				break
			}
		} else {
			break
		}
		pLast = pSlow
		pSlow = pSlow.next
	}

	pLast.next = pSlow.next // del
}

func main() {
	head := new(LinkedListNode)
	head.value = 0
	last := head
	mid := 4
	midNode := new(LinkedListNode)
	for i := 1; i < 9; i++ {
		node := new(LinkedListNode)
		node.value = i
		last.next = node
		last = node
		if i == mid {
			midNode = node
		}
	}
	for e := head; nil != e; e = e.next {
		fmt.Print(e.value, "-")
	}
	fmt.Println()
	// DeleteMid(head)
	DelNode(midNode)
	for e := head; nil != e; e = e.next {
		fmt.Print(e.value, "-")
	}
	fmt.Println()



}