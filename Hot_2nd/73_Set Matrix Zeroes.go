package hot2nd

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rowZero, colZero := false, false

	// Check if the first column has zeros
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colZero = true
		}
	}
	// Check if the first row has zeros
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			rowZero = true
		}
	}

	// // Record whether other rows and columns need to be set to zero
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// // According to the marking result, set the elements that need to be set to zero to zero
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// // If the first row needs to be set to zero, set all elements in the first row to zero
	if rowZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

	if colZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
