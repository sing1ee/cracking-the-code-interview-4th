package main

import (
	"fmt"
	"container/list"
)

type Pos struct {
	x, y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Queues8(queues []*Pos, index, row int) {
	if index == len(queues) {
		for _, v := range queues {
			fmt.Print("(", v.x, v.y,  ")")
		}
		fmt.Println()
		return
	}
	// fmt.Println(index)
	oklist := list.New();
	for col := 0; col < 8; col++ {
		ok := true
		for j := 0; j < index; j++ {
			if col == queues[j].y || abs(col - queues[j].y) == abs(row - queues[j].x) {
				ok = false
				break
			}
		}
		if ok {
			p := new(Pos)
			p.x = row
			p.y = col
			oklist.PushBack(p)
		}
	}
	for e := oklist.Front(); e != nil; e = e.Next() {
		queues[index] = e.Value.(*Pos)
		Queues8(queues, index + 1, row + 1)
	}
	return
}

func EightQueues(queues []*Pos) {
	for i := 0; i < 8; i++ {
		p := new(Pos)
		p.x = 0
		p.y = i
		queues[0] = p 
		Queues8(queues, 1, 1)
	}
}

func main() {
	queues := make([]*Pos, 8)
	EightQueues(queues)
}