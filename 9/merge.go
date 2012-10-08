package main

import (
	"fmt"
)

func Merge(a, b []int, la, lb int) {
	k := la + lb - 1
	i := la - 1
	j := lb - 1

	for ; j >= 0 && i >= 0; k--{
		if a[i] > b[j] {
			a[k] = a[i]
			i = i - 1
		} else {
			a[k] = b[j]
			j = j - 1
		}
	}

	for ; j >= 0; {
		a[k] = b[j]
		k = k -1
		j = j - 1
	}
}

func main() {
	a := []int{1, 3, 5, 7, 0, 0, 0}
	b := []int{2, 4, 6}
	Merge(a, b, 4, 3)
	fmt.Println(a)
}