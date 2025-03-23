package hot2nd

func spiralOrder(matrix [][]int) []int {
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	res := []int{}

	for top <= bottom && left <= right {
		//upper
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		//right
		if len(res) < len(matrix)*len(matrix[0]) {
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		//bottom
		if len(res) < len(matrix)*len(matrix[0]) {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		//left
		if len(res) < len(matrix)*len(matrix[0]) {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
