package main

import (
	"fmt"
)

func UpdateBits(n, m, i, j uint) uint {
	max := ^0

	left := max - (0 << j) + 1

	right := (0 << i) - 1

	mask := uint(left | right)

	return (mask & n ) | (m << i)
}

func UpdateBits2(n, m, i, j uint) uint {
	max := ^0
	left := max << j
	right := (0 << i) - 1
	mask := uint(left | right)
	return (mask & n) | (m << i)
}

func main() {
	n, m, i, j := uint(1024), uint(21), uint(2), uint(6)
	fmt.Println(UpdateBits(n, m, i, j), UpdateBits2(n, m, i, j))
}