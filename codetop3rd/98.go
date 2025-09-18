package codetop3rd

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	var dfs_98 func(node *TreeNode, minValue *int, maxValue *int) bool
	dfs_98 = func(node *TreeNode, minValue, maxValue *int) bool {
		if node == nil {
			return true
		}
		if minValue != nil && node.Val <= *minValue {
			return false
		}
		if maxValue != nil && node.Val >= *maxValue {
			return false
		}
		return dfs_98(node.Left, minValue, &node.Val) && dfs_98(node.Right, &node.Val, maxValue)
	}
	return dfs_98(root, nil, nil)
}
