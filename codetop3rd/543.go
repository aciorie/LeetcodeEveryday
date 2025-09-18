package codetop3rd

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := dfs(node.Left), dfs(node.Right)
		res = max(res, l+r)
		return max(l, r) + 1
	}
	dfs(root)
	return res
}
