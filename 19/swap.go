package main

import (
	"fmt"
)

func Swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func Swap2(a, b int) (int, int) {
	a = b - a
	b = b - a
	a = a + b
	return a, b
}

func main() {
	a, b := 1, 2
	fmt.Println(Swap2(a, b))
	fmt.Println(Swap(a, b))
}