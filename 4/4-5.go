package main

import (
	"fmt"
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
		l := root
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

