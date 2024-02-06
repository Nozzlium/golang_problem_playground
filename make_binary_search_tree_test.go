package main

import (
	"fmt"
	"testing"
)

func sortedArrayToBST(nums []int) *TreeNode {
	root := &TreeNode{}
	for _, num := range nums {
		root = appendRoot(root, num)
		fmt.Println(root.Val)
	}
	return root
}

func appendRoot(root *TreeNode, num int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: num,
		}
	}
	val := root.Val
	if num < val {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val: num,
			}
		} else {
			return &TreeNode{
				Val:   num,
				Right: root,
			}
		}
	}
	if num > val {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: num,
			}
		} else {
			return &TreeNode{
				Val:  num,
				Left: root,
			}
		}
	}
	return root
}

func TestMakeTree(t *testing.T) {
	sortedArrayToBST([]int{-10, -3, 0, 5, 9})
}
