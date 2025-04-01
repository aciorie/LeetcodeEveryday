package hot2nd

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n
	for l <= r {
		m := l + (r-l)/2
		if target == matrix[m/n][m%n] {
			return true
		} else if target > matrix[m/n][m%n] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}
