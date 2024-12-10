package hot

/*
Given an integer array nums of unique elements, return all possible
subsets
 (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]


Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.
*/

func Subsets(nums []int) [][]int {
	n, res := len(nums), make([][]int, 0)

	var dfs func(start int, current []int)
	dfs = func(start int, current []int) {
		// Append a copy of the current subset to the result
		subset := make([]int, len(current))
		copy(subset, current)
		res = append(res, subset)

		// Explore further elements to include in the subset
		for i := start; i < n; i++ {
			// Include nums[i] in the current subset
			current = append(current, nums[i])
			// Recur with the next index
			dfs(i+1, current)
			// Backtrack: remove the last element added
			current = current[:len(current)-1]
		}
	}

	// Start the DFS with an empty subset
	dfs(0, []int{})
	return res
}
