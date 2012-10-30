package main

import (
	"fmt"
)

func ContinuousLs(arr []int) int {
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	max := dp[0]
	for i := 1; i < len(arr); i++ {
		if dp[i - 1] + arr[i] > arr[i] {
			dp[i] = dp[i - 1] + arr[i]
		} else {
			dp[i] = arr[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	fmt.Println(dp)
	return max
}

func main() {
	arr := []int{0, 0, 0, 0, 0}
	fmt.Println(ContinuousLs(arr))
}