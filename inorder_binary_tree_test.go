package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverese(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, traverese(root.Left)...)
	res = append(res, root.Val)
	res = append(res, traverese(root.Right)...)
	return res
}

func inorderTraversal(root *TreeNode) []int {
	return traverese(root)
}
