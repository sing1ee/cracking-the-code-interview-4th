package main

import (
	"fmt"
)

type TreeNode struct {
	left, right *TreeNode
	value int
}

func sumStart(start *TreeNode, path []int, index, now, sum int)  {
	if nil == start {
		return
	}

	path[index] = start.value
	now = now + path[index]
	if now == sum {
		for i := 0; i <= index; i++ {
			fmt.Print(path[i], "-")
		}
		fmt.Println()
	}
	sumStart(start.left, path, index + 1, now, sum)
	sumStart(start.right, path, index + 1, now, sum)
}

func FindPaths(root *TreeNode, path []int, sum int) {
	if nil == root {
		return
	}
	sumStart(root, path, 0, 0, sum)
	FindPaths(root.left, path, sum)
	FindPaths(root.right, path, sum)
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
	FindPaths(root, path, 5)
}