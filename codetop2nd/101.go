package codetop2nd

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	return ismirror(root.Left, root.Right)
}

func ismirror(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if (l == nil && r != nil) || (l != nil && r == nil) {
		return false
	}
	if (l != nil && r != nil) && l.Val != r.Val {
		return false
	}
	return ismirror(l.Left, r.Right) && ismirror(l.Right, r.Left)
}
