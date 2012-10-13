package main

import (
	"fmt"
	"math/rand"
)

type LinkedListNode struct {
	value int
	next *LinkedListNode
}
func DeleteDups(head *LinkedListNode) {
	if nil == head {
		return
	}
	previous := head
	current := previous.next
	for ; nil != current; {
		runner := head
		for ; runner != current; {
			if runner.value == current.value {
				tmp := current.next
				previous.next = tmp
				current = tmp
				break
			}
			runner = runner.next
		}
		if runner == current {
			previous = current
			current = current.next
		}
	}
}

func main() {
	head := new(LinkedListNode)
	head.value = rand.Intn(10)
	last := head
	for i := 0; i < 10; i++ {
		node := new(LinkedListNode)
		node.value = rand.Intn(10)
		last.next = node
		last = node
	}
	for e := head; e != nil; e = e.next {
		fmt.Print(e.value, "-")
	}
	fmt.Println()
	DeleteDups(head)
	for e:= head; e != nil; e = e.next {
		fmt.Print(e.value, "-")
	}
	fmt.Println()
}


