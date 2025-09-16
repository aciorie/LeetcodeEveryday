package codetop3rd

func maxDepth(root *TreeNode) int {
	var dfs func(node *TreeNode, res int) int
	dfs = func(node *TreeNode, res int) int {
		if node == nil {
			return res
		}
		l, r := dfs(node.Left, res+1), dfs(node.Right, res+1)
		return max(l, r)
	}
	return dfs(root, 0)
}
