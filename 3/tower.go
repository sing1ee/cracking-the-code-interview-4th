package main

import (
	"fmt"
	"stack"
)

type Tower struct {
	disks *Stack
	index int
}

func (t *Tower) Index() int {
	return t.index
}

func (t *Tower) Add(d int) {
	if !disks.IsEmpty() && disks.Peek().(int) <= d {
		fmt.Println("error")
	} else {
		disks.Push(d)
	}
}

func (t *Tower) MoveTopTo(ot *Tower) {
	top := t.disks.Pop().(int)
	ot.Add(top)
	fmt.Println("Move disk", top, "from", t.Index(), "to", ot.Index())
}

func (t *Tower) Print() {
	fmt.Println("Contents of Tower", t.Index())
	for e := t.disks.Elements().Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

func (t *Tower) MoveDisks(n int, dest, buffer *Tower) {
	if n > 0 {
		t.MoveDisks(n - 1, buffer, dest)
		t.MoveTopTo(dest)
		buffer.MoveDisks(n - 1, dest, t)
	}
}

func main() {
	n := 5
	towers := make([]*Tower, 5)
	for i := 0; i < 5; i++ {
		towers[i] = &Tower{new(Stack), i}
	}
	for i := n - 1; i >= 0; i-- {
		towers[0].Add(i)
	}
	towers[0].MoveDisks(n, towers[2], towers[1])
}



