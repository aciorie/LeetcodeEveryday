package hot2nd

func rotate(matrix [][]int) {
	n := len(matrix)
	// row -> col, col -> row
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// inverse each row
	for i := 0; i < n; i++ {
		reverse_48(matrix[i])
	}
}

func reverse_48(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
