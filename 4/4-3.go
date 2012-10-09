package main 

import (
	"fmt"
)

type TreeNode struct {
	left, right *TreeNode
	value int
}

func BuildBT(arr []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	root := new(TreeNode)
	root.value = arr[mid]
	root.left = BuildBT(arr, left, mid - 1)
	root.right = BuildBT(arr, mid + 1, right)
	return root
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	depth := MaxDepth(root.left)
	if depth > max {
		max = depth
	}
	depth = MaxDepth(root.right)
	if depth > max {
		max = depth
	}
	return max + 1
}

func  MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := 100
	depth := MaxDepth(root.left)
	if depth < min {
		min = depth
	}
	depth = MaxDepth(root.right)
	if depth < min {
		min = depth
	}
	return min + 1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	root := BuildBT(arr, 0, len(arr) - 1)
	fmt.Println(MaxDepth(root), MinDepth(root))
}