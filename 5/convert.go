package main

import (
	"fmt"
)

func Convert(a, b int) int {
	count := 0

	for c := a ^ b; c > 0; c = c >> 1 {
		count = count + (c & 1)
	}
	return count
}

func main() {
	var a, b int
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println(Convert(a, b))
}