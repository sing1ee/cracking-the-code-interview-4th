package main 

import (
	"fmt"
	"container/list"
)

type TreeNode struct {
	left, right *TreeNode
	value int
}

func BuildBST(arr []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	root := new(TreeNode)
	root.value = arr[mid]
	root.left = BuildBST(arr, left, mid - 1)
	root.right = BuildBST(arr, mid + 1, right)
	return root
}

func LevelTraversal(root *TreeNode, size int) []*list.List {
	queue := make([]*TreeNode, size << 1)
	head := 0
	tail := 0
	queue[tail] = root
	tail++
	flag := new(TreeNode)
	flag.value = -1
	queue[tail] = flag
	tail++
	ll := make([]*list.List, size << 1)
	level := 0
	tmp := list.New()
	lastFlag := false
	for ; ; {
		el := queue[head]
		head++
		if el.value == -1 {
			if lastFlag {
				return ll
			}
			ll[level] = tmp
			level++
			tmp = list.New()
			queue[tail] = flag
			tail++
			lastFlag = true
		} else {
			lastFlag = false
			tmp.PushBack(el)
			if el.left != nil {
				queue[tail] = el.left
				tail++
			}
			if el.right != nil {
				queue[tail] = el.right
				tail++
			}
		}
	} 
	return ll
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	root := BuildBST(arr, 0, len(arr) - 1)
	ll := LevelTraversal(root, len(arr))
	for _, l := range ll {
		if l == nil {
			continue
		}
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value.(*TreeNode).value, " ")
		}
		fmt.Println()
	}
}