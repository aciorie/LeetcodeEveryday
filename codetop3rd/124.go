package codetop3rd

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var dfs_124 func(root *TreeNode) int
	dfs_124 = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lsum, rsum := max(0, dfs_124(root.Left)), max(0, dfs_124(root.Right))
		cursum := root.Val + lsum + rsum
		res = max(res, res+cursum)
		return max(lsum, rsum) + 1
	}
	dfs_124(root)
	return res
}
