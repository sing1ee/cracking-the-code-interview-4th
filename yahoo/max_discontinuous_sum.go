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

func MaxDisConSum2(arr []int) int {
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	dp[1] = arr[1]
	for i := 2; i < len(arr); i++ {
		if dp[i - 2] + arr[i] > dp[i - 1] {
			dp[i] = dp[i - 2] + arr[i]
		} else {
			dp[i] = dp[i - 1]
		}
	}
	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	arr := []int{6, 5, 4, 3, 2, 6, 10, 1, 5}
	fmt.Println(MaxDisConSum(arr), MaxDisConSum2(arr)) // ans: 2
}