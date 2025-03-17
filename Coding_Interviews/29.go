package codinginterviews

/*
面试题 29. 顺时针打印矩阵


题目描述
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


限制：

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	row, col := len(matrix), len(matrix[0])
	res := make([]int, row*col)
	top, bottom, left, right := 0, row-1, 0, col-1

	for top <= bottom && left <= right {
		// up side:from left to right
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// right side:from up to down
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// Check if there are any remaining rows/columns
		if top <= bottom && left <= right {
			// bottom side:from right to left
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--

			// left side:from bottom to top
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows, cols := len(matrix), len(matrix[0])
	res := make([]int, 0, rows*cols)
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// right, bottom, left, top
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIndex, row, col := 0, 0, 0

	for i := 0; i < rows*cols; i++ {
		res = append(res, matrix[row][col])
		visited[row][col] = true

		nextRow, nextCol := row+directions[dirIndex][0], col+directions[dirIndex][1]

		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols || visited[nextRow][nextCol] {
			dirIndex = (dirIndex + 1) % 4 //	Change directions
		}

		row += directions[dirIndex][0]
		col += directions[dirIndex][1]
	}

	return res
}
