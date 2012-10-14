package main

import (
	"fmt"
)

func Anagrams(s1, s2 string) bool {
	counter := make([]int, 26)
	a := int('a')
	for _, c := range s1 {
		index := int(c) - a
		counter[index]++
	}
	for _, c := range s2 {
		index := int(c) - a
		counter[index]--
		if counter[index] < 0 {
			return false
		}
	}
	for _, c := range counter {
		if c > 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "abceff"
	s2 := "abcefd"
	fmt.Println(Anagrams(s1, s2))
}
