package main

import (
	"fmt"
)

func NumOfZeros(n int) int {
	
	nums := 0

	for i := 5; n / i > 0; i = i * 5 {
		nums = nums + n / i
	}
	return nums
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(NumOfZeros(n))
}