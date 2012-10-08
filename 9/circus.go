package main

import (
	"fmt"
	"container/list"
	"math/rand"
)

type Tier struct {
	ht, wt int
}

func Quicksort(arr []*Tier, left, right int) {
	if left >= right {
		return
	}
	n := right - left
	index := left + rand.Intn(n)
	arr[left], arr[index] = arr[index], arr[left]
	small, large := left + 1, right
	for ; ; {
		for ; small <= right && arr[small].ht < arr[left].ht; small++ {}
		for ; arr[large].ht > arr[left].ht; large-- {}
		if small > large {
			break
		}
		arr[small], arr[large] = arr[large], arr[small]
	}
	arr[left], arr[large] = arr[large], arr[left]
	
	Quicksort(arr, left, large - 1)
	Quicksort(arr, large + 1, right)
}

func LIS(arr []*Tier) *list.List {
	dp := make([]int, len(arr))
	dp[0] = 1
	max, idx := 0, -1
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i].wt > arr[j].wt && dp[j] + 1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > max {
			max = dp[i]
			idx = i
		}
	}
	ls := list.New()
	for i := idx; i >= 0; i-- {
		if dp[i] == max {
			ls.PushBack(arr[i])
			max = max - 1
		} 
	}
	return ls
}

func main() {
	arr := []*Tier{&Tier{65, 100}, &Tier{70, 150}, &Tier{56, 90}, &Tier{75, 190}, &Tier{60, 95}, &Tier{68, 110}}
	Quicksort(arr, 0, len(arr) - 1)
	for _, v := range arr {
		fmt.Print(v.ht, " ")
	}
	fmt.Println()
	r := LIS(arr)
	for e := r.Back(); e != nil; e = e.Prev() {
		t := e.Value.(*Tier)
		fmt.Print("(", t.ht, t.wt, ")")
	}
	fmt.Println()
 }