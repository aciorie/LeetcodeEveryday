package matrix

func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]         // 左下角 -> 左上角
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j] // 右下角 -> 左下角
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i] // 右上角 -> 右下角
			matrix[j][n-1-i] = tmp                  // 左上角 -> 右上角
		}
	}
}
