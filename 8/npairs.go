package main

import (
	"fmt"
)
func Npairs(l, r, index int,  pairs []string) {
	if l < 0 || r < l {
		return
	}

	if l == 0 && r == 0 {
		fmt.Println(pairs)
	} else {
		if l > 0 {
			pairs[index] = "("
			Npairs(l - 1, r, index + 1, pairs)
		}

		if r > l {
			pairs[index] = ")"
			Npairs(l, r - 1, index + 1, pairs)
		}
	}
}

func main() {
	pairs := make([]string, 6)
	Npairs(3, 3, 0, pairs)
}