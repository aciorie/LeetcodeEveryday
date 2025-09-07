package codetop3rd

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs_200(grid, i, j)
			}
		}
	}
	return res
}

func dfs_200(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs_200(grid, i-1, j)
	dfs_200(grid, i+1, j)
	dfs_200(grid, i, j-1)
	dfs_200(grid, i, j+1)
}
