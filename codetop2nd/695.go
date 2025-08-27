package codetop2nd

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cur := dfs_695(grid, i, j)
				res = max(res, cur)
			}
		}
	}
	return res
}

func dfs_695(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	res := 1
	res += dfs_695(grid, i+1, j)
	res += dfs_695(grid, i-1, j)
	res += dfs_695(grid, i, j+1)
	res += dfs_695(grid, i, j-1)
	return res
}
