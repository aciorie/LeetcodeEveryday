package codetop

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	maxArea := 0

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				area := dfs_695(grid, r, c)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func dfs_695(grid [][]int, r, c int) int {
	m, n := len(grid), len(grid[0])
	if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0
	area := 1
	area += dfs_695(grid, r+1, c)
	area += dfs_695(grid, r-1, c)
	area += dfs_695(grid, r, c+1)
	area += dfs_695(grid, r, c-1)

	return area
}
