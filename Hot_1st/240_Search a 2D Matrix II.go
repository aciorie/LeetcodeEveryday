package hot

/*
Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.


Example 1:
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
Output: true

Example 2:
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
Output: false


Constraints:
m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-109 <= matrix[i][j] <= 109
All the integers in each row are sorted in ascending order.
All the integers in each column are sorted in ascending order.
-109 <= target <= 109
*/

//got tle but dont know why
func SearchMatrix_240(matrix [][]int, target int) bool {
	m := len(matrix)
	for i := 0; i < m; i++ {
		if matrix[i][0] > target {
			return false
		}
		if colBinarySearch(matrix[i], target) {
			continue
		} else {
			return true
		}
	}
	return false
}

func colBinarySearch(arr []int, target int) bool {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		if arr[m] == target {
			return true
		} else if arr[m] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	return false
}
