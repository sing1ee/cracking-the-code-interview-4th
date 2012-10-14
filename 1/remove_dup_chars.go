package main

import (
	"fmt"
)

func RemoveDups(arr []int) {
	last := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < last; j++ {
			if arr[i] == arr[j] {
				arr[j], arr[last] = arr[last], arr[j]
				last--
				j--
			}
		}
	}
	if last != len(arr) - 1 {
		for i := last; i < len(arr); i++ {
			arr[i] = -1
		}
	}
	
	fmt.Println(last, arr)
}

// no dups
func Case1() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	RemoveDups(arr)
}
// all dups
func Case2() {
	arr := []int{1, 1, 1, 1, 1, 1, 1}
	RemoveDups(arr)
}
// continous dups
func Case3() {
	arr := []int{1, 1, 1, 1, 2, 2, 2}
	RemoveDups(arr)
}
// discontinous dups
func Case4() {
	arr := []int{1, 2, 1, 2, 1, 2, 1, 2}
	RemoveDups(arr)
}

func main() {
	Case1()
	Case2()
	Case3()
	Case4()
}