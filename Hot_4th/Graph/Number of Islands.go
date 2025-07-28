package graph

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	res := 0
	var dfs_200 func(row, col int)
	dfs_200 = func(row, col int) {
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == '0' {
			return
		}

		grid[row][col] = '0'
		dfs_200(row+1, col)
		dfs_200(row-1, col)
		dfs_200(row, col+1)
		dfs_200(row, col-1)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '1' {
				res++
				dfs_200(r, c)
			}
		}
	}
	return res
}
