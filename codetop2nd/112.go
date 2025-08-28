package codetop2nd

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if targetSum == 0 {
		return true
	}
	var dfs_112 func(root *TreeNode, targetSum int) bool
	dfs_112 = func(root *TreeNode, targetSum int) bool {
		if root == nil {
			return false
		}
		targetSum -= root.Val
		if root.Left == nil && root.Right == nil {
			return targetSum == 0
		}
		return dfs_112(root.Left, targetSum) || dfs_112(root.Right, targetSum)
	}
	return dfs_112(root, targetSum)
}
