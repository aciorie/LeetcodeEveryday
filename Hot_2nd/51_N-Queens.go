package hot2nd

func solveNQueens(n int) [][]string {
	var res [][]string
	board := make([]string, n)
	for i := range board {
		board[i] = string(make([]rune, n))
		for j := range board[i] {
			board[i] = board[i][:j] + "." + board[i][j+1:]
		}
	}
	dfs_51(n, 0, &res, board, make([]bool, n), make([]bool, 2*n), make([]bool, 2*n))
	return res
}

func dfs_51(n int, row int, res *[][]string, board []string, columns []bool, diagonals1, diagonals2 []bool) {
	if row == n {
		solution := make([]string, n)
		copy(solution, board)
		*res = append(*res, solution)
		return
	}

	for col := 0; col < n; col++ {
		// If there is already a queen in that column or diagonal, skip it
		if columns[col] || diagonals1[row-col+n] || diagonals2[row+col] {
			continue
		}

		// Place the queens
		board[row] = board[row][:col] + "Q" + board[row][col+1:]
		columns[col] = true
		diagonals1[row-col+n] = true
		diagonals2[row+col] = true
		dfs_51(n, row+1, res, board, columns, diagonals1, diagonals2)
		board[row] = board[row][:col] + "." + board[row][col+1:]

		// Withdraw the move
		columns[col] = false
		diagonals1[row-col+n] = false
		diagonals2[row+col] = false
	}
}
