package codetop3rd

func spiralOrder(matrix [][]int) []int {
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	res := []int{}

	for top <= bottom && left <= right {
		// upper: left to right
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// right: top to bottom
		if top <= bottom && left <= right {
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		// bottom: right to left
		if top <= bottom && left <= right {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		// left: bottom to top
		if top <= bottom && left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
