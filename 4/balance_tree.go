package main

import (
	"fmt"
	"container/list"
)

type TreeNode struct {
	Children *list.List
	value int
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	for e := root.Children.Front(); e != nil; e = e.Next() {
		m := MaxDepth(e.Value.(*TreeNode))
		if m > depth {
			depth = m
		}
	}
	return depth + 1
}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Children.Front() == nil {
		return 1
	}
	depth := 100
	for e := root.Children.Front(); e != nil; e = e.Next() {
		m := MinDepth(e.Value.(*TreeNode))
		if m < depth {
			depth = m
		}
	}
	return depth + 1
}

func IsBalanced(root *TreeNode) bool {
	maxDepth, minDepth := MaxDepth(root), MinDepth(root)
	fmt.Println(maxDepth, minDepth)
	return maxDepth == minDepth || minDepth + 1 == maxDepth
}

func main() {
	root := &TreeNode{list.New(), 0}
	for i := 1; i < 5; i++ {
		node := &TreeNode{list.New(), i}
		root.Children.PushBack(node)
	}
	for i := 5; i < 10; i++ {
		node := &TreeNode{list.New(), i}
		root.Children.Front().Value.(*TreeNode).Children.PushBack(node)
	}
	for i := 10; i < 15; i++ {
		node := &TreeNode{list.New(), i}
		root.Children.Front().Value.(*TreeNode).Children.Front().Value.(*TreeNode).Children.PushBack(node)	
	}
	fmt.Println(IsBalanced(root))
}