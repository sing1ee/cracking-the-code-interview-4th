package main

import (
	"fmt"
)

/*
非连续子数组的最大和
*/
func MaxDisConSum(arr []int) int {
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	dp[1] = arr[1]
	max := 0
	for i := 2; i < len(arr); i++ {
		for j := 0; j < i - 1; j++ {
			if dp[i] < dp[j] + arr[i] {
				dp[i] = dp[j] + arr[i]
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Print(dp[i], " ")
	}
	fmt.Println()
	return max
}

func main() {
	arr := []int{1, 1, 1}
	fmt.Println(MaxDisConSum(arr))
}