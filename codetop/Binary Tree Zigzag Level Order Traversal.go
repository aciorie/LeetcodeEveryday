package codetop

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := []*TreeNode{root}
	isLeftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		curLevel := make([]int, 0)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if isLeftToRight {
				curLevel = append(curLevel, node.Val)
			} else {
				curLevel = append([]int{node.Val}, curLevel...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, curLevel)
		isLeftToRight = !isLeftToRight
	}

	return res
}
