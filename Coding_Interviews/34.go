package codinginterviews

/*
面试题 34. 二叉树中和为某一值的路径


题目描述
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。



示例 1：



输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]
示例 2：



输入：root = [1,2,3], targetSum = 5
输出：[]
示例 3：

输入：root = [1,2], targetSum = 0
输出：[]


提示：

树中节点总数在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000
*/
func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var curPath []int
	var dfs func(node *TreeNode, curSum int)
	dfs = func(node *TreeNode, curSum int) {
		if node == nil {
			return
		}

		// Add the value of the current node to the path
		curPath = append(curPath, node.Val)
		curSum += node.Val

		// Check if it is a leaf node
		if node.Left == nil && node.Right == nil {
			if curSum == target {
				// Add a copy of the current path to the result
				tmp := make([]int, len(curPath))
				copy(tmp, curPath)
				res = append(res, tmp)
			}
		}

		// Recursively traverse the left and right subtrees
		dfs(node.Left, curSum)
		dfs(node.Right, curSum)

		// Backtrack and remove the value of the current node
		curPath = curPath[:len(curPath)-1]
	}

	dfs(root, 0)
	return res
}
