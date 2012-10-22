package main

import (
	"fmt"
)

type TreeNode struct {
	left, right *TreeNode
	value int
}

func (t *TreeNode) IsBalance() bool {

	maxDepth := t.MaxDepth()
	minDepth := t.MinDepth()
	if maxDepth == minDepth || minDepth + 1 == maxDepth {
		return true
	}
	return false
}

func (t *TreeNode) MaxDepth() int {
	if nil == t {
		return 0
	}
	leftDepth := t.left.MaxDepth()
	rightDepth := t.right.MaxDepth()

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func (t *TreeNode) MinDepth() int {
	if nil == t {
		return 0
	}
	leftDepth := t.left.MinDepth()
	rightDepth := t.right.MinDepth()

	if leftDepth < rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}


func BuildBalanceTree(arr []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	root := new(TreeNode)
	root.left = BuildBalanceTree(arr, left, mid - 1)
	root.right = BuildBalanceTree(arr, mid + 1, right)
	root.value = arr[mid]
	return root
}

func NormalCase() {
	root := new(TreeNode)
	root.value = 0
	fmt.Println(root.IsBalance())
	left1 := new(TreeNode)
	left1.value = 1
	root.left = left1
	fmt.Println(root.IsBalance())
	left2 := new(TreeNode)
	left2.value = 2
	left1.left = left2
	fmt.Println(root.IsBalance())
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	root := BuildBalanceTree(arr, 0, len(arr) - 1)
	fmt.Println(root.IsBalance())
	NormalCase()
}



