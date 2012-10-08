package main

import (
	"fmt"
)

func Ncents(n, denom int) int {
	nextDenom := 0
	switch denom {
	case 25:
		nextDenom = 10
		break
	case 10:
		nextDenom = 5
		break
	case 5:
		nextDenom = 1
		break
	case 1:
		return 1
	}
	ways := 0;
	for i := 0; i * denom <= n; i++ {
		ways += Ncents(n - i * denom, nextDenom)
	}
	return ways
}

func main() {
	fmt.Println(Ncents(100, 25))
}