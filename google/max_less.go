package main

import (
	"fmt"
)
/*
#Google面试题#给一个排序好的整数数组A，请写一个函数，输入是数组A和一个整数x，返回数组A中值小于x的最大元素的索引和值。
假设 升序
*/
func BinarySearch(arr []int, x int) (int, int) {
	left, right := 0, len(arr) - 1
	
	for ; left <= right; {
		mid := (left + right) >> 1
		if arr[mid] == x {
			if mid == 0 {
				return -1, -1
			}
			return mid - 1, arr[mid - 1]
		}

		if arr[mid] > x {
			right = mid - 1
		} else {
			if mid + 1 == len(arr) {
				return mid, arr[mid]
			}
			if arr[mid + 1] >= x {
				return mid, arr[mid]
			} else if arr[mid + 1] < x {
				left = mid + 1
			}
		}
	}
	return -1, -1
}

//=
func Case1() {
	arr := []int{1, 2, 3, 4, 5}
	for _, v := range arr {
		fmt.Println(BinarySearch(arr, v))
	}
}
func Case2() {
	arr := []int{1, 3, 5, 7, 9}
	fmt.Println(BinarySearch(arr, 2))
	fmt.Println(BinarySearch(arr, 4))
	fmt.Println(BinarySearch(arr, 6))
	fmt.Println(BinarySearch(arr, 8))
	fmt.Println(BinarySearch(arr, 10))
}
func main() {
	// Case1()	
	// Case2()
}