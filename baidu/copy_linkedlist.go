package main

import (
	"fmt"
)

type LinkedListNode struct {
	value int

	pNext, pSibling *LinkedListNode
}

func (lln *LinkedListNode) Clone() *LinkedListNode {
	head := new(LinkedListNode)
	head.value = lln.value


	for next := lln.pNext; nil != next; {
		head.pNext = next
		next = next.pNext
	}

	oldlln := lln
	newlln := head
	for ; nil != old; {
		oldtmp := oldlln.pNext
		oldlln.pNext = newlln
		newtmp := newlln.pNext
		newlln.pNext = oldtmp

		oldlln = oldtmp
		newlln = newtmp
	}

	for start := lln; nil != start; {
		next := start.pNext
		next.pSibling = start.pSibling
		start = next.pNext
	}

	oldStart := lln
	newStart := head
	for ; nil != oldStart; {
		oldStart.pNext = newStart.next
		oldStart = newStart.next
		newStart.pNext = oldStart.pNext
		newStart = oldStart.pNext
	}
	return head
}

func main() {
	
}