package main

import (
	"fmt"
	"time"
	"math/rand"
)

func Rand15() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(5) + 1
}

func Rand17() int {
	ret := 0
	for ; ; {
		ret = 5 * (Rand15() - 1) + Rand15()
		if ret <= 21 {
			break;
		}
	}
	return 1 + ret % 7
}

func main() {
	arr := make([]int, 5)
	for i := 0; i < 2000; i++ {
		arr[Rand15() - 1]++
	}
	fmt.Println(arr)
}