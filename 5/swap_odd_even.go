package main

import (
	"fmt"
)

func Swap(a int) int {
	
	return ((a & 0x55555555) << 1) | ((a & 0xAAAAAAAA) >> 1)
}

func main() {
	var a int
	fmt.Scanf("%d", &a)
	fmt.Println(Swap(a))
}