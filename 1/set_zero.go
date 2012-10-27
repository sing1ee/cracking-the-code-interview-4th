package main

import (
	"fmt"
)

func SetZero(matrix [][]int) {
	rows := make([]int, len(matrix))
	cols := make([]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if 0 == matrix[i][j] {
				rows[i] = 1
				cols[j] = 1s
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if 1 == cols[j] || 1 == rows[i] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	
}