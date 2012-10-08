package main

import (
	"fmt"
	"container/list"
)

func Permu(str []int, index int) *list.List {
	l := len(str)
	tmp := []int{ str[index] }
	if index + 1 == l {
		r := list.New()
		r.PushFront(tmp)
		return r
	}
	ps := list.New()
	r := Permu(str, index + 1)
	for e := r.Front(); e != nil; e = e.Next(){
		p := e.Value.([]int)
		for i := 0; i <= len(p); i++ {
			before := make([]int, i - 0)
			copy(before, p[0:i])
			after := make([]int, len(p) - i)
			copy(after, p[i:])
			ps.PushFront(append(append(before, str[index]), after...))
		}
	}
	return ps
}

func main() {
	str := []int{1, 2, 3, 4}
	l := Permu(str, 0)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}