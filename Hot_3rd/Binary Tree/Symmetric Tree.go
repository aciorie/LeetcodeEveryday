package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if (l == nil && r != nil) || (l != nil && r == nil) {
		return false
	}
	if (l != nil && r != nil) && l.Val != r.Val {
		return false
	}
	return isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
}
