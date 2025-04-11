package hot2nd

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		level := queue.Len()
		var cur []int

		for i := 0; i < level; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			cur = append(cur, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, cur)
	}
	return res
}
