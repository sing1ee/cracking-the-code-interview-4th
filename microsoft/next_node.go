package main

import (
	"fmt"
)

type TreeNode struct {
	value int
	left, right, father *TreeNode
}

func Next(t *TreeNode) *TreeNode {
	if nil != t.right {

		left := t.right
		for ; ; {
			if nil == left.left {
				break
			}
			left = left.left
		}
		return left
	} else {
		parent := t.father
		if t == parent.left {
			return parent
		} else {
			for ; nil != parent; {
				if parent.father.left == parent {
					return parent
				}
				parent = parent.father
			} 
		}
		return parent
	}
	return nil
}

func BuildBST(arr []int, left, right int, father *TreeNode)  *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right - left) >> 1
	root := new(TreeNode)
	root.value = arr[mid]
	root.father = father
	root.left = BuildBST(arr, left, mid - 1, root)
	root.right = BuildBST(arr, mid + 1, right, root)
	return root
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	root := BuildBST(arr, 0, len(arr) - 1, nil)

	fmt.Println(root.value, Next(root).value)
}




