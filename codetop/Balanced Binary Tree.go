package codetop

func isBalanced(root *TreeNode) bool {
	return checkBalance(root) != -1
}

func checkBalance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh, rh := checkBalance(root.Left), checkBalance(root.Right)
	if lh == -1 || rh == -1 {
		return -1
	}
	if abs(lh, rh) > 1 {
		return -1
	}

	return max(lh, rh) + 1
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
