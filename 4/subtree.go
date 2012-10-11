package main

import (
	"fmt"
)

type TreeNode struct {
	left, right *TreeNode
	value int
}

func matchTree(leftRoot, rightRoot *TreeNode) bool {

	if nil == rightRoot {
		return true
	}

	if nil == leftRoot || leftRoot.value != rightRoot.value {
		return false
	}

	return matchTree(leftRoot.left, rightRoot.left) && matchTree(leftRoot.right, rightRoot.right)
}

func SubTree(bigRoot, smallRoot * TreeNode) bool {
	if nil == bigRoot {
		return false
	}	
	if nil == smallRoot {
		return true
	}
	if matchTree(bigRoot, smallRoot) {
		return true
	}
	return SubTree(bigRoot.left, smallRoot) || SubTree(bigRoot.right, smallRoot)
}

func BuildTree(arr []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	root := new(TreeNode)
	root.value = arr[mid]
	root.left = BuildTree(arr, left, mid - 1)
	root.right = BuildTree(arr, mid + 1, right)
	return root
}

func Case1() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bigRoot := BuildTree(arr1, 0, len(arr1) - 1)
	arr2 := []int{2, 5, 7}
	smallRoot := BuildTree(arr2, 0, len(arr2) - 1)
	fmt.Println(SubTree(bigRoot, smallRoot))
}

func Case2() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bigRoot := BuildTree(arr1, 0, len(arr1) - 1)
	arr2 := []int{1, 2, 5, 7}
	smallRoot := BuildTree(arr2, 0, len(arr2) - 1)
	fmt.Println(SubTree(bigRoot, smallRoot))
}

func main() {
	Case1()	
	Case2()
}
