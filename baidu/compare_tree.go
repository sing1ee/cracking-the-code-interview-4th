package main

import (
	"fmt"
)

type TreeNode struct {
	left, right *TreeNode
	value int
}

func (t *TreeNode) Equals(other *TreeNode) bool {
	if nil == t && nil == other {
		return true
	}

	if nil == t || nil == other {
		return false
	}

	if t.value != other.value {
		return false
	}

	return (t.left.Equals(other.left) && t.right.Equals(other.right)) || (t.left.Equals(other.right) && t.right.Equals(other.left))
}

func BuildTree(arr []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right - left) >> 1
	root := new(TreeNode)
	root.value = arr[mid]
	root.left = BuildTree(arr, left, mid - 1)
	root.right = BuildTree(arr, mid + 1, right)
	return root
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	arr2 := []int{7, 6, 5, 4, 3, 2, 1}
	arr3 := []int{7, 6, 5, 4, 3, 2, 1}
	arr4 := []int{1, 2, 3, 4, 5, 6}
	root1 := BuildTree(arr1, 0, len(arr1) - 1)
	root2 := BuildTree(arr2, 0, len(arr2) - 1)
	root3 := BuildTree(arr3, 0, len(arr3) - 1)
	root4 := BuildTree(arr4, 0, len(arr4) - 1)
	fmt.Println(root1.Equals(root2))
	fmt.Println(root1.Equals(root3))
	fmt.Println(root1.Equals(root4))
}