package main

import (
	"fmt"
	"strings"
	"strconv"
)

func Float2Binary(floatStr string) string {
	idx := strings.Index(floatStr, ".")
	if -1 == idx {
		idx = len(floatStr)
	}

	intPart, ei := strconv.Atoi(floatStr[0:idx])
	floatPart, ef := strconv.ParseFloat("0" + floatStr[idx:], 64)
	if nil != ei || nil != ef {
		return "ERROR"
	}

	intPartStr := ""
	floatPartStr := ""
	for ; intPart > 0; {
		m := intPart % 2
		intPart = intPart >> 1
		intPartStr = strconv.Itoa(m) + intPartStr
	}
	for ; floatPart > 0; {
		if len(floatPartStr) > 32 {
			return "ERROR"
		}

		r := floatPart * 2

		if r > 1 {
			floatPartStr = floatPartStr + "1"
			floatPart = floatPart - 1
		} else {
			floatPartStr = floatPartStr + "0"
			floatPart = r
		}
	}
	return intPartStr + "." + floatPartStr
}

func main() {
	fmt.Println(Float2Binary("12.5"))
}