package codinginterviews

/*
面试题 33. 二叉搜索树的后序遍历序列


题目描述
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true


提示：

数组长度 <= 1000
*/
func verifyPostorder(postorder []int) bool {
	var dfs func(l, r int) bool
	dfs = func(l, r int) bool {
		// Current subarray is empty or has only one element
		if l >= r {
			return true
		}
		v, i := postorder[r], l

		// Find the starting position of the right subtree
		for i < r && postorder[i] < v {
			i++
		}

		// Verify that all nodes in the right subtree are greater than the root node value v
		for j := i; j < r; j++ {
			if postorder[j] < v {
				return false
			}
		}
		// Recursively verify the left and right subtrees
		return dfs(l, i-1) && dfs(i, r-1)
	}
	return dfs(0, len(postorder)-1)
}
