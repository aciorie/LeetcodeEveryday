package Dec

import (
	"container/list"
	"sort"
)

/*
515. Find Largest Value in Each Tree Row
Medium
Topics
Companies
Given the root of a binary tree, return an array of the largest value in each row of the tree (0-indexed).



Example 1:
Input: root = [1,3,2,5,3,null,9]
Output: [1,3,9]

Example 2:
Input: root = [1,2,3]
Output: [1,3]


Constraints:

The number of nodes in the tree will be in the range [0, 104].
-231 <= Node.val <= 231 - 1
*/

func LargestValues(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		var level []int

		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			level = append(level, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		sort.Ints(level)
		res = append(res, level[len(level)-1])
	}

	return res
}
