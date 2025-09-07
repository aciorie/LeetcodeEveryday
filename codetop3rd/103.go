package codetop3rd

import (
	"container/list"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	fromLeftToRight := true
	for queue.Len() > 0 {
		levelSize := queue.Len()
		var cur []int
		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			cur = append(cur, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		if !fromLeftToRight {
			for i, j := 0, len(cur)-1; i < j; i, j = i+1, j-1 {
				cur[i], cur[j] = cur[j], cur[i]
			}
		}
		res = append(res, cur)
		fromLeftToRight = !fromLeftToRight
	}
	return res
}
