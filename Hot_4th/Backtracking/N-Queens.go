package backtracking

func solveNQueens(n int) [][]string {
	var res [][]string

	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	var dfs_51 func(row int, board [][]byte, cols, diag1, diag2 []bool, n int, res *[][]string)
	dfs_51 = func(row int, board [][]byte, cols, diag1, diag2 []bool, n int, res *[][]string) {
		if row == n {
			ans := make([]string, n)
			for i := 0; i < n; i++ {
				ans[i] = string(board[i])
			}
			*res = append(*res, ans)
			return
		}

		for col := 0; col < n; col++ {
			d1Index, d2Index := row-col+n-1, row+col
			if !cols[col] && !diag1[d1Index] && !diag1[d2Index] {
				board[row][col] = 'Q'
				cols[col] = true
				diag1[d1Index] = true
				diag2[d2Index] = true

				dfs_51(row+1, board, cols, diag1, diag2, n, res)

				board[row][col] = '.'
				cols[col] = false
				diag1[d1Index] = false
				diag2[d2Index] = false
			}
		}
	}
	return res
}
