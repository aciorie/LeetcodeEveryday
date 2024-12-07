package hot

/*
You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.



Example 1:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Example 2:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false


Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104
*/

/*
Wrong Answer 1
132 / 133 testcases passed

Input matrix = [[-8,-7,-5,-3,-3,-1,1],[2,2,2,3,3,5,7],[8,9,11,11,13,15,17],[18,18,18,20,20,20,21],[23,24,26,26,26,27,27],[28,29,29,30,32,32,34]]
target = -5

Output
false
Expected
true

cause: in "find in columns" part, my setting was	"for lCol < rCol"
*/

/*
Wrong Answer 2
Time Limit Exceeded
2 / 133 testcases passed


Use Testcase
matrix = [[1]]
target = 0
*/
func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	l, r := 0, m-1
	row := -1

	//check rows first because it said next row[0]>prev row[n-1]
	for l <= r {
		mid := l + (r-l)/2

		//find the exact target
		if matrix[mid][0] == target || target == matrix[mid][n-1] {
			return true
		}

		if matrix[mid][0] < target && target < matrix[mid][n-1] {
			row = mid
			break
		} else if matrix[mid][0] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if row == -1 {
		return false
	}

	//find in columns
	lCol, rCol := 0, n-1
	for lCol <= rCol {
		mid := lCol + (rCol-lCol)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			lCol = mid + 1
		} else {
			rCol = mid - 1
		}
	}
	return false
}

/*
Wrong Answer 3	107 / 133 testcases passed

Input
matrix = [[1]]
target = 1

Output	false
Expected	true
*/
func SearchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix)-1, len(matrix[0])-1
	l, r := 0, m*n
	for l < r {
		m := l + (r-l)/2
		if target == matrix[m/n][m%n] {
			return true
		} else if target > matrix[m/n][m%n] {
			l = m + 1
		} else {
			r = m
		}
	}
	return false
}

//correct one
func SearchMatrix3(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
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
