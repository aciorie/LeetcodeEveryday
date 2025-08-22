package codetop2nd

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var dfs_94 func(node *TreeNode)
	dfs_94 = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs_94(node.Left)
		res = append(res, node.Val)
		dfs_94(node.Right)
	}
	dfs_94(root)
	return res
}
