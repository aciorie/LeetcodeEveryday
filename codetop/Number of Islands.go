package codetop

func numIslands(grid [][]byte) int {
	res := 0
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				res++               // 岛屿数量加一
				dfs_200(grid, r, c) // 从这个陆地开始，淹没整个岛屿
			}
		}
	}
	return res
}

func dfs_200(grid [][]byte, r, c int) {
	rows := len(grid)
	cols := len(grid[0])

	// 边界条件：
	// 1. 坐标越界   2. 当前单元格是水 ('0')
	if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
		return
	}
	// 将当前陆地标记为水 ('0')，表示已访问，避免重复计数和无限递归
	grid[r][c] = '0'

	dfs_200(grid, r+1, c) //down
	dfs_200(grid, r-1, c) //up
	dfs_200(grid, r, c+1) //right
	dfs_200(grid, r, c-1) //left
}
