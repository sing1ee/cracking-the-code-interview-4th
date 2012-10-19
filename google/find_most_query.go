package main

/*
#Google面试题#给你一天的Google搜索日志，你怎么设计算法找出是否有一个搜索词，
它出现的频率占所有搜索的一半以上？如果肯定有一个搜索词占大多数，你能怎么提高你的算法找到它？
再假定搜索日志就是内存中的一个数组，能否有O(1)空间，O(n)时间的算法？

1）我的思路：O（N）中位数得思想
2）假定第一个查询是最多的，从头开始遍历，候选设为第一个，数量是1，处理第2个，如果与候选相同，则数量++
	如果不同，则数量--，当数量为0时，更新候选，并且把数量变成1
*/

import (
	"fmt"
)

func Find2(arr []string) string {
	if len(arr) == 0 {
		return "no"
	}
	can := arr[0]
	num := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] == can {
			num++
		} else {
			num--
		}
		if 0 == num {
			can = arr[i]
			num = 1
		}
	}
	return can
}

func main() {
	arr := []string{"a", "b", "c", "b", "b", "a", "b"}
	fmt.Println(Find2(arr))
}