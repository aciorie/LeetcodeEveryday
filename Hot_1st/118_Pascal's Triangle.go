package hot

/*
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:




Example 1:
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

Example 2:
Input: numRows = 1
Output: [[1]]


Constraints:

1 <= numRows <= 30
*/
func Generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, row)
	}
	return res
}
