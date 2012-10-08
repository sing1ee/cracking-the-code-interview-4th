package main

import (
	"fmt"
)

func Subsets(set []int, index int) [][]int {
	l := len(set)
	if index == l - 1 {
		result := make([][]int, 2)
		result[0] = []int{}
		result[1] = []int{ set[index] }
		return result
	}

	tmp := Subsets(set, index + 1)
	result := make([][]int, len(tmp) * 2)
	for i := 0; i < len(tmp); i++ {
		result[i] = append(result[i], tmp[i]...)
	}
	for i := len(tmp); i < len(tmp) * 2; i++ {
		tmp[i - len(tmp)] = append(tmp[i - len(tmp)], set[index])
		result[i] = append(result[i], tmp[i - len(tmp)]...)
	}
	return result
}

func pow (a , b int) int {
	r := 1
	for i := 0; i < b; i++ {
		r = r * a
	}
	return r
}

func Sub(set []int) [][]int {
	l := len(set)
	n := pow(2, l)
	result := make([][]int, n)
	index := 0
	for i := 0; i < n; i++ {
		idx := 0
		for k:= i ; k > 0; {
			if k & 1 == 1 {
				result[index] = append(result[index], set[idx])
			}		
			k = k >> 1
			idx = idx + 1
		}
		index = index + 1
	}
	return result
}

func main() {
	arr := []int{1, 2, 3}
	result := Subsets(arr, 0)
	for _, ss := range result {
		fmt.Println("+", ss)
	}
	result = Sub(arr)
	for _, ss := range result {
		fmt.Println("-", ss)
	}	
}