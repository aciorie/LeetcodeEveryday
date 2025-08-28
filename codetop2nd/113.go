package codetop2nd

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	cur := []int{}
	var dfs_113 func(root *TreeNode, targetSum int)

	dfs_113 = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		
		cur = append(cur, root.Val)
		if root.Left == nil && root.Right == nil && root.Val == targetSum {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		}

		dfs_113(root.Left, targetSum-root.Val)
		dfs_113(root.Right, targetSum-root.Val)

		cur = cur[:len(cur)-1]
	}

	dfs_113(root, targetSum)
	return res
}
