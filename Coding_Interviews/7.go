package codinginterviews

/*
面试题 07. 重建二叉树


题目描述
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



示例 1:



Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
示例 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]


限制：

0 <= 节点个数 <= 5000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solutio7_1(preorder []int, inorder []int) *TreeNode {
	d := make(map[int]int)
	for i, v := range inorder {
		d[v] = i
	}

	// i and j to represent the starting positions of the preorder and inorder sequences, and n to represent the number of nodes
	var dfs func(i, j, n int) *TreeNode
	dfs = func(i, j, n int) *TreeNode {
		if n == 0 {
			return nil
		}
		// root position
		k := d[preorder[i]]
		// number of left subtree nodes
		l := k - j
		root := &TreeNode{Val: preorder[i]}
		// left subtree
		root.Left = dfs(i+1, j, l)
		// right subtree
		root.Right = dfs(i+1+l, k+1, n-l-1)
		return root
	}
	return dfs(0, 0, len(inorder))
}
