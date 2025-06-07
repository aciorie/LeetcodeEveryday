package review

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if dfs_79(board, word, 0, row, col, visited) {
				return true
			}
		}
	}
	return false
}

func dfs_79(board [][]byte, word string, index int, row int, col int, visited [][]bool) bool {
	if len(word) == index {
		return true
	}
	m, n := len(board), len(board[0])
	if row < 0 || row >= m || // 行越界
		col < 0 || col >= n || // 列越界
		board[row][col] != word[index] || // 当前字符不匹配
		visited[row][col] { // 当前单元格已在当前路径中被访问过
		return false
	}

	visited[row][col] = true
	// 定义方向数组：上、下、左、右
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // {row_change, col_change}
	for _, dir := range directions {
		nextRow, nextCol := row+dir[0], col+dir[1]
		if dfs_79(board, word, index+1, nextRow, nextCol, visited) {
			return true
		}
	}

	visited[row][col] = false
	return false
}
