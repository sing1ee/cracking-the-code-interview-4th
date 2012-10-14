package main

import (
	"fmt"
)
/*
1)if there is additional space
2)if the chars are all ascii
3)...
*/
func IsUniqueChar(str string) bool {
	checker := 0
	start := int('a')
	for _, c := range str {
		val := uint(int(c) - start)
		if checker & (1 << val) > 0 {
			return false
		}
		checker = checker | (1 << val)
	}
	return true
}

func main() {
	str := "abcdef"
	fmt.Println(IsUniqueChar(str))
}