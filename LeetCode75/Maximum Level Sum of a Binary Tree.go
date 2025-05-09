package leetcode75

import "container/list"

func maxLevelSum(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	curLevel := 1
	maxLevel, maxLevelSum := 1, root.Val
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		curLevelSum := 0
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			curLevelSum += node.Val

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		if curLevelSum > maxLevelSum {
			maxLevel = curLevel
			maxLevelSum = curLevelSum
		}
		curLevel++
	}
	return maxLevel
}
