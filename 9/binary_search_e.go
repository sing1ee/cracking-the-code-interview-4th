package main 

import (
	"fmt"
)

func EmptyBinarySearh(arr [] string, left, right int, t string) int {
	for ; left <= right; {
		mid := (left + right) >> 1
		if arr[mid] == t {
			return mid
		}
		idx := mid
		for ; arr[idx] == "" && idx >= 0; idx-- {}
		if idx >= 0{
			if t <= arr[idx] {
				right = idx
			} else {
				left = mid + 1
			}
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	fmt.Println(arr)
	for i, v := range arr {
		if v == "" {
			continue
		}
		fmt.Println(i, v, EmptyBinarySearh(arr, 0, len(arr) - 1, v))
	}

	fmt.Println(EmptyBinarySearh(arr, 0, len(arr) - 1, "zhang"))
}