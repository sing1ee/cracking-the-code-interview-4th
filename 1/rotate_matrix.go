package main

import (
	"fmt"
)

func RotateMatrix(matrix [][]int) {
	
	for level := 0; level < len(matrix) / 2; level++ {

		first, last := level, len(matrix) - 1 - level
		for i := first; i < last; i++ {
			top := matrix[level][i]
			offset := i - first
			matrix[level][i] = matrix[last - offset][level]
			matrix[last - offset][level] = matrix[last][last - offset]
			matrix[last][last - offset] = matrix[i][last]
			matrix[i][last] = top
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	for _, v := range matrix {
		fmt.Println(v)
	}
	fmt.Println("after rotate")
	RotateMatrix(matrix)

	for _, v := range matrix {
		fmt.Println(v)
	}
}