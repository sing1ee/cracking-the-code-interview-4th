package main

import (
	"fmt"
)

func AddWithoutOper(a, b int) int {
	
	if 0 == b {
		return a
	}

	sum := a ^ b
	carry := (a & b) << 1

	return AddWithoutOper(sum, carry)
}

func main() {
	fmt.Println(AddWithoutOper(1, 29))
}