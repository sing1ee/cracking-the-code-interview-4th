package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
// return 1 with prob p = 0.7
func Randp() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(10)
	if n <= 6 {
		return 1
	} 
	return 0
}

// return 1 with prob p = 0.5

func Rande() int {
	a, b := Randp(), Randp()
	if 1 == a && 0 == b {
		return 1
	} else if 0 == a && 1 == b {
		return 0
	} 
	return Rande()
}

func Randn(n int) int {
	k := 1 + int(math.Log2(float64(n)))
	result := 0
	for i := 0; i < k; i++ {
		if 1 == Rande() {
			result = result | (1 << uint(i))
		}
	}
	if result >= n {
		return Randn(n)
	}
	return result
}


func main() {
	n := 3
	arr := make([]int, n)
	for i := 0; i < 2000; i++ {
		arr[Randn(n)]++
	}
	fmt.Println(arr)
}