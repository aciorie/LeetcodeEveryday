package hot

/*
Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

Example 1:
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.

Example 2:
Input: heights = [2,4]
Output: 4


Constraints:
1 <= heights.length <= 105
0 <= heights[i] <= 104
*/

//first attempting and failed. i checked some disscussions but didn't get them.
//some guys say you can consider the samller one from both side,but how to recognize one's value super big?
//if i follow this thought i can think of binary search,but like e.g. 1, you can't check this cases
//anyways,i'll checktheir solutions.	maybe next time i'll write them in a .md file lol
func largestRectangleArea(heights []int) int {
	n, res := len(heights), 0
	if n == 1 {
		res = heights[0]
	}
	return res
}
