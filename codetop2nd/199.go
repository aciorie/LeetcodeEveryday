package codetop2nd

import "container/list"

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)

			if i == levelSize-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}
