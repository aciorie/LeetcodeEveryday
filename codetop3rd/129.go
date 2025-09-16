package codetop3rd

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, sum int) int
	dfs = func(node *TreeNode, sum int) int {
		if node == nil {
			return sum
		}
		sum = sum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		return dfs(node.Left, sum) + dfs(node.Right, sum)
	}
	return dfs(root, 0)
}
