package review

func solveNQueens(n int) [][]string {
	var res [][]string

	// 初始化棋盘
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	// 初始化辅助数组
	cols := make([]bool, n) // 标记列是否被占用

	// 每个皇后必须在不同的主对角线（行索引 - 列索引相等）和副对角线（行索引 + 列索引相等）
	diag1 := make([]bool, 2*n-1) // 标记主对角线 (r - c) 是否被占用, 范围从 -(n-1) 到 (n-1)，映射到 0 到 2*n-2
	diag2 := make([]bool, 2*n-1) // 标记副对角线 (r + c) 是否被占用, 范围从 0 到 2*n-2

	backtrack_51(0, board, cols, diag1, diag2, n, &res)
	return res
}

// 回溯函数
// row: 当前正在处理的行
// board: 当前棋盘状态
// cols, diag1, diag2: 标记数组
// n: 棋盘大小
// result: 存储所有解决方案的指针
func backtrack_51(row int, board [][]byte, cols, diag1, diag2 []bool, n int, res *[][]string) {
	// 基本情况
	if row == n {
		solution := make([]string, n)
		for i := 0; i < n; i++ {
			solution[i] = string(board[i])
		}
		*res = append(*res, solution)
		return
	}

	// 遍历当前行的所有列
	for col := 0; col < n; col++ {
		// 计算当前位置对应的对角线索引
		// 主对角线: r - c + offset (offset 通常是 n-1，将负值映射到非负索引)
		// 副对角线: r + c
		d1Index := row - col + n - 1
		d2Index := row + col

		// 剪枝/合法性检查
		// 检查当前列、主对角线、副对角线是否已被占用
		if !cols[col] && !diag1[d1Index] && !diag2[d2Index] {
			// 放置皇后
			board[row][col] = 'Q'
			cols[col] = true
			diag1[d1Index] = true
			diag2[d2Index] = true

			// 递归：尝试在下一行放置皇后
			backtrack_51(row+1, board, cols, diag1, diag2, n, res)

			// 回溯/撤销：移除皇后，为其他可能性做准备
			board[row][col] = '.'
			cols[col] = false
			diag1[d1Index] = false
			diag2[d2Index] = false
		}
	}
}
