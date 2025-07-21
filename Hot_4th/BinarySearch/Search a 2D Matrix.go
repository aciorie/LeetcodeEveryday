package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return false
	}

	lowRow, highRow := 0, n-1
	targetRow := -1
	for lowRow <= highRow {
		midRow := lowRow + (highRow-lowRow)/2
		if matrix[midRow][0] == target {
			return true
		}
		if matrix[midRow][0] < target {
			if matrix[midRow][n-1] >= target {
				targetRow = midRow
				break
			}
			lowRow = midRow + 1
		} else {
			highRow = midRow - 1
		}
	}
	if targetRow == -1 {
		return false
	}

	rowToSearch := matrix[targetRow]
	lowCol, highCol := 0, n-1

	for lowCol <= highCol {
		midCol := lowCol + (highCol-lowCol)/2
		if rowToSearch[midCol] == target {
			return true
		} else if rowToSearch[midCol] < target {
			lowCol = midCol + 1
		} else {
			highCol = midCol - 1
		}
	}

	return false
}
