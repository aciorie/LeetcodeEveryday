package contest

import (
	"fmt"
	"sort"
)

/*
Q1. Sort Matrix by Diagonals
Medium
3 pt.
You are given an n x n square matrix of integers grid. Return the matrix such that:

The diagonals in the bottom-left triangle (including the middle diagonal) are sorted in non-increasing order.
The diagonals in the top-right triangle are sorted in non-decreasing order.


Example 1:

Input: grid = [[1,7,3],[9,8,2],[4,5,6]]

Output: [[8,2,3],[9,6,7],[4,5,1]]

Explanation:



The diagonals with a black arrow (bottom-left triangle) should be sorted in non-increasing order:

[1, 8, 6] becomes [8, 6, 1].
[9, 5] and [4] remain unchanged.
The diagonals with a blue arrow (top-right triangle) should be sorted in non-decreasing order:

[7, 2] becomes [2, 7].
[3] remains unchanged.
Example 2:

Input: grid = [[0,1],[1,2]]

Output: [[2,1],[1,0]]

Explanation:



The diagonals with a black arrow must be non-increasing, so [0, 2] is changed to [2, 0]. The other diagonals are already in the correct order.

Example 3:

Input: grid = [[1]]

Output: [[1]]

Explanation:

Diagonals with exactly one element are already in order, so no changes are needed.



Constraints:

grid.length == grid[i].length == n
1 <= n <= 10
-105 <= grid[i][j] <= 105
*/

/*
Wrong Answer
114 / 1306 testcases passed
Input
grid =
[[1,7,3],[9,8,2],[4,5,6]]
Use Testcase
Output
[[6,9,8],[7,4,2],[3,5,1]]
Expected
[[8,2,3],[9,6,7],[4,5,1]]
*/
func SortMatrix(grid [][]int) [][]int {
	if len(grid) == 1 {
		return [][]int{{1}}
	}
	n := len(grid)

	fmt.Println("Original grid:", grid)

	// bottom left
	for col := 0; col < n; col++ {
		// get diagonals
		diag := []int{}
		x, y := 0, col
		for x < n && y >= 0 {
			diag = append(diag, grid[x][y])
			x++
			y--
		}

		// sort
		sort.Sort(sort.Reverse(sort.IntSlice(diag)))
		x, y = 0, col
		for i := 0; i < len(diag); i++ {
			grid[x][y] = diag[i]
			x++
			y--
		}
	}

	// top right
	for row := 1; row < n; row++ {
		diag := []int{}
		x, y := row, n-1
		for x < n && y >= 0 {
			diag = append(diag, grid[x][y])
			x++
			y--
		}
		sort.Ints(diag)
		x, y = row, n-1
		for i := 0; i < len(diag); i++ {
			grid[x][y] = diag[i]
			x++
			y--
		}
	}

	diag := []int{}
	for i := 0; i < n; i++ {
		diag = append(diag, grid[i][i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(diag)))
	for i := 0; i < n; i++ {
		grid[i][i] = diag[i]
	}

	return grid
}
