package main

import (
	"fmt"
)

func SubMul(arr []int, out int) int {
	
	a := make([]int, len(arr))
	a[0] = 1
	for i := 1; i < len(arr); i++ {
		a[i] = a[i - 1] * arr[i - 1] 
	}
	b := make([]int, len(arr))
	b[len(arr) - 1] = 1
	for i := len(arr) - 2; i >= 0; i-- {
		b[i] = b[i + 1] * arr[i + 1]
	}

	return a[out] * b[out]
}

func main() {
	arr := []int{1, 2, 3, 4}
	for v, _ := range arr {
		fmt.Println(SubMul(arr, v))
	}
}