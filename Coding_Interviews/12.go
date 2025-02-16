package codinginterviews

/*
面试题 12. 矩阵中的路径


题目描述
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。





示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false


提示：

1 <= board.length <= 200
1 <= board[i].length <= 200
board 和 word 仅由大小写英文字母组成
*/

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// If the current character word[k] has matched the end of the string word, it means that the path to the string word has been found.
		if k == len(word) {
			return true
		}

		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
			return false
		}

		// Mark the current character as visited
		board[i][j] = ' '
		// Up, Right, Down, Left
		dirs := []int{-1, 0, 1, 0, -1}
		res := false

		for l := 0; l < 4; l++ {
			res = res || dfs(i+dirs[l], j+dirs[l+1], k+1)
		}
		board[i][j] = word[k]
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
