package main

import (
	"fmt"
	"math"
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

func Cls(arr []int) int {
	sumMax := int(math.MinInt32)
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if sum > sumMax {
			sumMax = sum
		} else {
			sum = 0
		}
	}
	return sumMax
}

func main() {
	arr := []int{1, 2, 3, -1}
	fmt.Println(ContinuousLs(arr), Cls(arr))
}