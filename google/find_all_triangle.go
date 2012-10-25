package main

/*
#Google面试题#给一个无序的正整数数组，找出所有三个元素的组合使它们作为三条边能形成一个三角形。
比如，输入为{4, 6, 3, 7}, 可能组合为 {3, 4, 6}，{4, 6, 7}和{3, 6, 7}。尽量优化你的算法。
*/

import (
	"fmt"
	"time"
	"math/rand"
)

func Qsort(arr[] int, left, right int) {
	if left >= right {
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := left + r.Intn(right - left)
	arr[left], arr[index] = arr[index], arr[left]
	i, j := left + 1, right
	for ; ; {
		for ; i <= right && arr[i] < arr[left]; {
			i++
		}
		for ; j >= 0 && arr[j] > arr[left]; {
			j--
		}
		if j < i {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[left], arr[j] = arr[j], arr[left]
	Qsort(arr, j + 1, right)
	Qsort(arr, left, j - 1)
}

func FindAllTria(arr []int) {
	Qsort(arr, 0, len(arr) - 1)
	for i := len(arr) - 1; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			if arr[j] > arr[j - 1] < arr[i] {
				break
			}
			for k := j - 1; k >= 0; k-- {
				if arr[k] + arr[j] <= arr[i] {
					break;
				}
				fmt.Println(arr[k], arr[j], arr[i])
			}
		}
	}
}

func main() {
	arr := []int{3, 1, 4, 2, 5}
	FindAllTria(arr)
}