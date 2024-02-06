package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		return q == nil
	}
	if q == nil {
		return p == nil
	}
	a := p.Val == q.Val
	b := isSameTree(p.Left, q.Left)
	c := isSameTree(p.Right, q.Right)
	return a && b && c
}
