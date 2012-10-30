package main

import (
	"fmt"
)

func MaxNoif(a, b int) int {
	
	c := a - b
	mark := c >> 31 & 0x1
	return a - mark * c
}

func main() {

	var a, b int
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println(MaxNoif(a, b))
}