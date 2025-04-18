package hot2nd

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32

	// Calculate the maximum path sum from the current node as the root node to the leaf node
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// Recursively calculate the contribution values ​​of left and right child nodes
		leftGain, rightGain := max(maxGain(node.Left), 0), max(maxGain(node.Right), 0)
		// Calculate the maximum path and
		priceNewPath := node.Val + leftGain + rightGain

		// Update the global maximum path and
		res = max(res, priceNewPath)
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return res
}
