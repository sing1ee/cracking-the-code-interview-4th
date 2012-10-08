package main

import (
	"fmt"
)

func Fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return Fib(n - 1) + Fib(n - 2)
}

func FibDp(n int) int {
	dp := make([]int, n + 1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}

func main() {
	n := 100

	fmt.Println(Fib(n) == FibDp(n))
}