package main

import (
	"fmt"
)
/**
*return the index of t if t is in the array, else -1
*/
func RotateBinarySearch(array []int, left, right, t int) int {
	for ; left <= right; {
		mid := (left + right) >> 1

		if t == array[mid] {
			return mid
		}
		if t > array[mid] {
			if t <= array[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if t >= array[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

func main() {
	arr := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	for i, v := range arr {
		fmt.Println(i, v, RotateBinarySearch(arr, 0, len(arr) - 1, v))
	}
	fmt.Println(RotateBinarySearch(arr, 0, len(arr) - 1, 100))
}