package main

import (
	"fmt"
	"math/rand"
)
/*
* Write an algorithm to find the 'next' node (in-order successor) of a given
* node in a binary search tree where each node has a link to its parent
*/

type TreeNode struct {
	left, right, parent *TreeNode
	value int
}

func FindNext(root, node *TreeNode) *TreeNode {
	if node.right != nil {
		l := root.right
		for ; l.left != nil; l = l.left { }
		return l
	} else {
		p := node.parent
		for ; p != nil; {
			if p.left == node {
				return p
			}
			node = p
			p = node.left
		}
	}
	return nil
}

func BuildBST(arr []int, left, right int, parent *TreeNode) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	root := new(TreeNode)
	root.parent = parent
	root.value = arr[mid]
	root.left = BuildBST(arr, left, mid - 1, root)
	root.right = BuildBST(arr, mid + 1, right, root)
	return root
}

func Case1() {

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	root := BuildBST(arr, 0, len(arr) - 1, nil)
	next := FindNext(root, root)
	fmt.Println(root.value, next.value)
}

func Case2() {
	root := new(TreeNode)
	root.value = 0
	parent := root
	for i := 1; i < 10; i++ {
		left := new(TreeNode)
		left.parent = parent
		left.value = i
		parent.left = left
		parent = left
	}
	next := FindNext(root, root)
	fmt.Println(root.value, next.value)
}

func Case3() {
	root := new(TreeNode)
	root.value = 0
	parent := root
	node := new(TreeNode)
	index := rand.Intn(9) + 1
	for i := 1; i < 10; i++ {
		left := new(TreeNode)
		left.parent = parent
		left.value = i
		parent.left = left
		if i == index {
			node = left
		}
		parent = left

	}
	next := FindNext(root, node)
	fmt.Println(root.value, next.value, index)
}

func main() {
	Case1()
	Case3()
	Case2()
	
}
