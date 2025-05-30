package codinginterviews

/*
面试题 26. 树的子结构


题目描述
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || (A == nil && B != nil) {
		return false
	}
	var dfs func(A, B *TreeNode) bool
	dfs = func(A, B *TreeNode) bool {
		if B == nil {
			return true
		}

		if A == nil || A.Val != B.Val {
			return false
		}

		return dfs(A.Left, B.Left) && dfs(A.Right, B.Right)
	}
	if B == nil {
		return false
	}
	if A == nil {
		return false
	}
	return dfs(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
