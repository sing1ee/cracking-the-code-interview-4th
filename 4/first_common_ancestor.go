package main

import (
	"fmt"
)

type TreeNode struct {
	left, right *TreeNode
	value int
}

func CommonAncestor(root, p, q *TreeNode) *TreeNode {
	if cover(root.left, p) && cover(root.left, q) {
		return CommonAncestor(root.left, p, q)
	}
	if cover(root.right, p) && cover(root.right, q) {
		return CommonAncestor(root.right, p, q)
	}
	return root
}

func cover(root, p *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p {
		return true
	}
	return cover(root.left, p) || cover(root.right, p)
}

func Case1() {
	root := new(TreeNode)
	root.value = 0
	left1 := new(TreeNode)
	left1.value = 1
	root.left = left1
	left2 := new(TreeNode)
	left2.value = 2
	left1.left = left2

	fmt.Println(CommonAncestor(root, left1, left2).value)
}

func Case2() {
	root := new(TreeNode)
	root.value = 0
	left1 := new(TreeNode)
	left1.value = 1
	root.left = left1
	left2 := new(TreeNode)
	left2.value = 2
	right2 := new(TreeNode)
	right2.value = 3
	left1.left = left2
	left1.right = right2
	fmt.Println(CommonAncestor(root, left2, right2).value)
}

func main() {
	Case1()
	Case2()
}