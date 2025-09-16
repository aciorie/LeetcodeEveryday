package codetop3rd

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	var dfs_101 func(l, r *TreeNode) bool
	dfs_101 = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if (l == nil && r != nil) || (l != nil && r == nil) {
			return false
		}
		if (l != nil && r != nil) && l.Val != r.Val {
			return false
		}
		return dfs_101(l.Left, r.Right) && dfs_101(l.Right, r.Left)
	}
	return dfs_101(root.Left, root.Right)
}
