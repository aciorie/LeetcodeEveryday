package matrix

// TLE
func rotate(matrix [][]int) {
	n := len(matrix)
	// row to col, col to row
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// inverse every row
	for i := 0; i < n; i++ {
		inverse_array(matrix[i])
	}
}

func inverse_array(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
