package codetop2nd

func sumNumbers(root *TreeNode) int {
	var dfs_129 func(node *TreeNode, curSum int) int
	dfs_129 = func(node *TreeNode, curSum int) int {
		if node == nil {
			return 0
		}

		curSum = curSum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return curSum
		}
		return dfs_129(node.Left, curSum) + dfs_129(node.Right, curSum)
	}
	return dfs_129(root, 0)
}
