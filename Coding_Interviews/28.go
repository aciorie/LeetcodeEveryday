package codinginterviews

/*
面试题 28. 对称的二叉树


题目描述
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3



示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false


限制：

0 <= 节点个数 <= 1000
*/
func isSymmetric(root *TreeNode) bool {
	var dfs func(a, b *TreeNode) bool
	dfs = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if (a == nil && b != nil) || (a != nil && b == nil) || a.Val != b.Val {
			return false
		}
		return dfs(a.Left, b.Right) && dfs(a.Right, b.Left)
	}
	return dfs(root, root)
}
