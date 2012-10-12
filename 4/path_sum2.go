package main

import (
	"fmt"
)

type TreeNode struct {
	left, right *TreeNode
	value int
}

func FindPaths(root *TreeNode, arr []int, index, sum int) {
	if nil == root {
		return
	}

	arr[index] = root.value
	tmp := sum
	for i := index; i >= 0; i-- {
		tmp = tmp - arr[i]
		
		if tmp < 0 {
			break
		}
		// fmt.Println(tmp, index, arr)
		if tmp == 0 {
			// find a path from a[0:index]
			for j := i; j <= index; j++ {
				fmt.Print(arr[j], "-")
			}
			fmt.Println()
		}
	}
	// fmt.Println("=====")
	FindPaths(root.left, arr, index + 1, sum)
	FindPaths(root.right, arr, index + 1, sum)
}

func BuildTestTree(arr []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	root := new(TreeNode)
	root.value = arr[mid]
	root.left = BuildTestTree(arr, left, mid - 1)
	root.right = BuildTestTree(arr, mid + 1, right)
	return root
} 

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := BuildTestTree(arr, 0, len(arr) - 1)
	path := make([]int, len(arr))
	FindPaths(root, path, 0, 13)
}