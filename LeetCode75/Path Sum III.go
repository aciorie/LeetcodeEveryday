package leetcode75

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	countAtRoot := countPathsFromNode(root, 0, targetSum)
	countInLeft := pathSum(root.Left, targetSum)
	countInRight := pathSum(root.Right, targetSum)

	return countAtRoot + countInLeft + countInRight
}

func countPathsFromNode(node *TreeNode, currentSum int, targetSum int) int {
	if node == nil {
		return 0
	}

	currentSum += node.Val
	count := 0
	if currentSum == targetSum {
		count = 1
	}
	count += countPathsFromNode(node.Left, currentSum, targetSum)
	count += countPathsFromNode(node.Right, currentSum, targetSum)

	return count
}
