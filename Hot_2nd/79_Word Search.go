package hot2nd

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs_79(board, word, 0, i, j, visited) {
				return true
			}
		}
	}
	return false
}

func dfs_79(board [][]byte, word string, index int, i int, j int, visited [][]bool) bool {
	if len(word) == index {
		return true
	}
	m, n := len(board), len(board[0])
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || word[index] != board[i][j] {
		return false
	}
	visited[i][j] = true
	res := dfs_79(board, word, index+1, i-1, j, visited) ||
		dfs_79(board, word, index+1, i+1, j, visited) ||
		dfs_79(board, word, index+1, i, j-1, visited) ||
		dfs_79(board, word, index+1, i, j+1, visited)

	visited[i][j] = false
	return res
}
