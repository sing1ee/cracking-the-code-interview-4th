package main

import (
	"fmt"
)

func ReplaceBlank(str []int, length int) {
	blankCount := 0;
	for i := 0; i < length; i++ {
		if str[i] == 0 {
			blankCount++
		}
	}
	newLen := length + (blankCount << 1)
	fmt.Println(length, newLen)
	lastIndex := newLen - 1
	for i := length - 1; i >= 0; i-- {
		if str[i] == 0 {
			str[lastIndex] = -3
			lastIndex--
			str[lastIndex] = -2
			lastIndex--
			str[lastIndex] = -1
			lastIndex--
		} else {
			str[lastIndex] = str[i]
			lastIndex--
		}
	}
}

func main() {
	arr := []int{1, 0, 2, 0, 0, 3, 0, -1, -1, -1, -1, -1, -1, -1, -1}
	ReplaceBlank(arr, 7)
	fmt.Println(arr)
}