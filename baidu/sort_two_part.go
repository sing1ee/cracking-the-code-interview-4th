package main

import (
	"fmt"
)

/*
百度面试题（一）：假设一整型数组存在若干正数和负数，现在通过某种算法使得该数组的所有负数在正数的左边，
且保证负数和正数间元素相对位置不变。时空复杂度要求分别为：o(n)和o(1)。

百度面试题（二），给定一个存放正数的数组，重新排列数组使得数组左边为奇数，右边为偶数，
且保证奇数和偶数之间元素相对位置不变。时空复杂度要求分别为：o(n)和o(1)。
*/

func PartSort(arr []int) {
	last := len(arr) - 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] & 1 == 0 {
			arr[i], arr[last] = arr[last], arr[i]
			last--
		}
	}
}

func main() {
	arr := []int{1, 2, 3}
	PartSort(arr)
	fmt.Println(arr)
}