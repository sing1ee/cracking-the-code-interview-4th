package main

import (
	"fmt"
	"time"
	"math/rand"
)

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t :=  left + r.Intn(right - left + 1)
	// fmt.Println(left, right, t, arr)
	arr[left], arr[t] = arr[t], arr[left]
	i, j := left + 1, right
	for ; ; {
		for ; i <= right && arr[i] <= arr[left]; { // this = is very important
			i++
		}
		for ; arr[j] > arr[left]; {
			j--
		}
		if j < i {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[left], arr[j] = arr[j], arr[left]
	// fmt.Println(arr, j)
	QuickSort(arr, left, j - 1)
	QuickSort(arr, j + 1, right)
}

func FindTwoSum(arr []int, sum int) {
	QuickSort(arr, 0, len(arr) - 1)
	first, last := 0, len(arr) - 1
	for ; first <= last; {
		tmp := arr[first] + arr[last]
		if tmp == sum {
			fmt.Println(arr[first], arr[last])
			first++
			last--
		} else if tmp > sum {
			last--
		} else {
			first++
		}
	}
}

func main() {
	
	arr := []int{1, 2, -1, -2, 5, 4, -10}
	FindTwoSum(arr, 3)
}