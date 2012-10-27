package main

import (
	"fmt"
	"strings"
)

func IsRotate(s1, s2 string) bool {
	str := s1 + s1
	return strings.Contains(str, s2)
}

func main() {
	s1 := "abcdef"
	s2 := "defabc"

	fmt.Println(IsRotate(s1, s2))
}