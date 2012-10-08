package main

import (
	"fmt"
)

func MatrixBinarySearch(arr [][]int, m, n, t int) (int, int) {
	leftRow, rightRow := 0, m - 1
	row := -1
	for ; leftRow <= rightRow; {
		midRow := (leftRow + rightRow) >> 1

		if arr[midRow][0] <= t && arr[midRow][n - 1] >= t {
			row = midRow
			break
		}
		if arr[midRow][0] <= t {
			leftRow = midRow + 1
		} else {
			rightRow = midRow - 1
		}
	}
	if row == -1 {
		return -1, -1
	}
	leftCol, rightCol := 0, n - 1
	for ; leftCol <= rightCol; {
		midCol := (leftCol + rightCol) >> 1

		if arr[row][midCol] == t {
			return row, midCol
		}
		if arr[row][midCol] > t {
			rightCol = midCol - 1
		} else {
			leftCol = midCol + 1
		}
	}
	return row, -1
}

func Find(matrix [][]int, m, n, t int) (int, int) {
	row, col := 0, n - 1
	for ; row < m && col >=0; {
		if (matrix[row][col] == t) {
			return row, col
		} else if (matrix[row][col] > t) {
			col = col - 1
		} else {
			row = row + 1
		}
	}
	return -1, -1
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for row, arr := range matrix {
		for col, v := range arr {
			x, y := Find(matrix, 3, 3, v)
			fmt.Println(row, col, v, x, y)
		}
	}
}