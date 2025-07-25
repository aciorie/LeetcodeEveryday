package binarytree

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var dfs_543 func(root *TreeNode) int

	dfs_543 = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := dfs_543(root.Left), dfs_543(root.Right)
		res = max(res, l+r)
		return max(l, r) + 1
	}
	dfs_543(root)
	return res
}
