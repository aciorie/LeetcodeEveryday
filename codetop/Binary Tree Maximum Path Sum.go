package codetop

import "math"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftGain, rightGain := dfs(root.Left), dfs(root.Right)
		curPathSum := root.Val + leftGain + rightGain
		res = max(res, curPathSum)

		return root.Val + max(leftGain, rightGain)
	}
	dfs(root)
	return res
}
