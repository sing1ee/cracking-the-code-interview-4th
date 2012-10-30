package main

import (
	"fmt"
)

func AddNoOper(a, b int) int {
	if 0 == b {
		return  a
	}
	c := a ^ b
	carry := a & b << 1
	return AddNoOper(c, carry)
}

func AddNoOper2(a, b int) int {
	return (a * a - b * b) / (a - b)
}

func main() {
	var a, b int
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println(AddNoOper(a, b), AddNoOper2(a, b))
}