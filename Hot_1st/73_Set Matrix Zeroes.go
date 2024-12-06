package hot

/*
Given an m x n integer matrix matrix, if an element is 0,
set its entire row and column to 0's.
You must do it in place.


Example 1:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]


Constraints:

m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1


Follow up:

A straightforward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/

//failed attempts, 61 / 202 testcases passed
func SetZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	isZero := make([][]bool, m)
	for i := range isZero {
		isZero[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isZero[i][j] {
				continue
			}
			if matrix[i][j] == 0 {
				makeZero(matrix, i, j)
			}
		}
	}
}

func makeZero(matrix [][]int, i, j int) {
	m, n := len(matrix), len(matrix[0])

	//make row all 0
	for a := 0; a < n; a++ {
		matrix[i][a] = 0
	}

	//make column all 0
	for b := 0; b < m; b++ {
		matrix[b][j] = 0
	}
}

//try with the solutions
func SetZeroes2(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRowZero, firstColZero := false, false

	// Check if the first row and first column need to be zero
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
		}
	}

	// Mark zeros in the first row and column
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Set matrix cells to zero based on markers
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Zero out the first row if needed
	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// Zero out the first column if needed
	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
