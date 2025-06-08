package review

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(0, dfs(node.Left))
		rightGain := max(0, dfs(node.Right))

		curParhSum := node.Val + leftGain + rightGain
		res = max(res, curParhSum)

		return node.Val + max(leftGain, rightGain)
	}

	dfs(root)
	return res
}
