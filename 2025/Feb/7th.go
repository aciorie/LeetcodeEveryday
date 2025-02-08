package feb

/*
3160. Find the Number of Distinct Colors Among the Balls
Medium
Topics
Companies
Hint
You are given an integer limit and a 2D array queries of size n x 2.

There are limit + 1 balls with distinct labels in the range [0, limit]. Initially, all balls are uncolored. For every query in queries that is of the form [x, y], you mark ball x with the color y. After each query, you need to find the number of distinct colors among the balls.

Return an array result of length n, where result[i] denotes the number of distinct colors after ith query.

Note that when answering a query, lack of a color will not be considered as a color.



Example 1:

Input: limit = 4, queries = [[1,4],[2,5],[1,3],[3,4]]

Output: [1,2,2,3]

Explanation:



After query 0, ball 1 has color 4.
After query 1, ball 1 has color 4, and ball 2 has color 5.
After query 2, ball 1 has color 3, and ball 2 has color 5.
After query 3, ball 1 has color 3, ball 2 has color 5, and ball 3 has color 4.
Example 2:

Input: limit = 4, queries = [[0,1],[1,2],[2,2],[3,4],[4,5]]

Output: [1,2,2,3,4]

Explanation:



After query 0, ball 0 has color 1.
After query 1, ball 0 has color 1, and ball 1 has color 2.
After query 2, ball 0 has color 1, and balls 1 and 2 have color 2.
After query 3, ball 0 has color 1, balls 1 and 2 have color 2, and ball 3 has color 4.
After query 4, ball 0 has color 1, balls 1 and 2 have color 2, ball 3 has color 4, and ball 4 has color 5.


Constraints:

1 <= limit <= 109
1 <= n == queries.length <= 105
queries[i].length == 2
0 <= queries[i][0] <= limit
1 <= queries[i][1] <= 109
*/

/*
Wrong Answer
213 / 551 testcases passed

Editorial
Input
limit =
1
queries =
[[0,1],[1,4],[1,1],[1,4],[1,1]]

Use Testcase
Output
[1,2,1,1,1]
Expected
[1,2,1,2,1]
*/
func QueryResults(limit int, queries [][]int) []int {
	// colors of balls
	ballMap := make(map[int]int)
	colorCounts := make(map[int]int)
	colorMap := make(map[int]struct{})
	res := make([]int, len(queries))

	for i, query := range queries {
		x, y := query[0], query[1]

		currentColor, exists := ballMap[x]
		if exists {
			colorCounts[currentColor]--
			if colorCounts[currentColor] == 0 {
				delete(colorCounts, currentColor)
				delete(colorMap, currentColor)
			}
		}

		ballMap[x] = y
		colorCounts[y]++
		if colorCounts[y] == 1 {
			colorMap[y] = struct{}{}
		}

		res[i] = len(colorMap)
	}
	return res
}
