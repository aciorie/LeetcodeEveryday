package codetop2nd

import "container/list"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	islefttoright := true
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		levelSize := queue.Len()
		var cur []int

		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if islefttoright {
				cur = append(cur, node.Val)
			} else {
				cur = append([]int{node.Val}, cur...)
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, cur)
		islefttoright = !islefttoright
	}
	return res
}
