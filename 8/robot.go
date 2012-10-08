package main

import (
	"fmt"
)

func Recur(n, x, y int) int {
	if x == 0 && y == 0 {
		return 0
	}
	if x == 1 || y == 1 {
		return 1
	}
	return Recur(n, x - 1, y) + Recur(n, x, y - 1)
}

func Dp(n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[0][i], dp[i][0] = 1, 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}
	return dp[n - 1][n - 1]
}

func C(n int) uint64 {
	var a, b uint64
	a, b = 1, 1

	var i uint64
	var l uint64
	l = uint64(n)
	for i = l - 1; i > 0; i-- {
		b = b * i
	}

	for i = 2 * l - 2; i > 0; i-- {

		a = a * i
	}

	return a / (b * b)
}

func main() {
	n := 10

	fmt.Println(Recur(n, n, n))
	fmt.Println(Dp(n))
	fmt.Println(C(n))
}