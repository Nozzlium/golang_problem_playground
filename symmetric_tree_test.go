package main

func isSymmetric(root *TreeNode) bool {
	left := inorderTraversalLeftRight(root.Left)
	right := inorderTraversalRightLeft(root.Right)
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func inorderTraversalLeftRight(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversalLeftRight(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversalLeftRight(root.Right)...)
	return res
}

func inorderTraversalRightLeft(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversalRightLeft(root.Right)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversalRightLeft(root.Left)...)
	return res
}
