package main

import (
	"fmt"
)


/*
*(1)if p is on the left of root, q is on the right, then root is the first common ancestor
*(2)if p and q are both one the left side, the deal with left sub tree recursively
*
*/


type TreeNode struct {
	left, right *TreeNode
	value int
}

const (
	NO_NODES_FOUND = iota
	ONE_NODES_FOUND
	TWO_NODES_FOUND
)

func covers(root, p, q *TreeNode) int {
	ret := NO_NODES_FOUND
	if root == nil {
		return ret
	}
	if root == p || root == q {
		ret += 1
	}
	ret += covers(root.left, p, q)
	if ret == TWO_NODES_FOUND {
		return ret
	}
	return ret + covers(root.right, p, q)
}

func CommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == q && (root.left == q || root.right == p) {
		return root
	}
	leftNum := covers(root.left, p, q)
	if leftNum == TWO_NODES_FOUND {
		if root.left == p || root.left == q {
			return root.left
		} else {
			return CommonAncestor(root.left, p, q)
		}
 	}
 	rightNum := covers(root.right, p, q)
 	if rightNum == TWO_NODES_FOUND {
 		if root.right == p || root.right == q {
 			return root.right
 		} else {
 			return CommonAncestor(root.right, p, q)
 		}
 	}
 	if rightNum == ONE_NODES_FOUND && leftNum == ONE_NODES_FOUND {
 		return root
 	}
 	return nil
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