package Dec

/*
An array is considered special if every pair of its adjacent elements contains two numbers with different parity.

You are given an array of integer nums and a 2D integer matrix queries, where for queries[i] = [fromi, toi] your task is to check that
subarray
 nums[fromi..toi] is special or not.

Return an array of booleans answer such that answer[i] is true if nums[fromi..toi] is special.



Example 1:
Input: nums = [3,4,1,2,6], queries = [[0,4]]
Output: [false]
Explanation:
The subarray is [3,4,1,2,6]. 2 and 6 are both even.

Example 2:
Input: nums = [4,3,1,6], queries = [[0,2],[2,3]]
Output: [false,true]
Explanation:
The subarray is [4,3,1]. 3 and 1 are both odd. So the answer to this query is false.
The subarray is [1,6]. There is only one pair: (1,6) and it contains numbers with different parity. So the answer to this query is true.


Constraints:

1 <= nums.length <= 105
1 <= nums[i] <= 105
1 <= queries.length <= 105
queries[i].length == 2
0 <= queries[i][0] <= queries[i][1] <= nums.length - 1
*/

//got TLE   535 / 536 testcases passed
func IsArraySpecial(nums []int, queries [][]int) []bool {
	res := make([]bool, len(queries))

	for i, query := range queries {
		from, to := query[0], query[1]
		isSpecial := true

		//check if subarray nums[from..to] is special
		for j := from; j < to; j++ {
			if (nums[j] % 2) == (nums[j+1] % 2) {
				//check parity of adjacent elements
				isSpecial = false
				break
			}
		}

		res[i] = isSpecial
	}

	return res
}

//one on the solution page
func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	// Step 1: Precompute prefix parity
	isSpecialPrefix := make([]bool, n)
	for i := 1; i < n; i++ {
		isSpecialPrefix[i] = (nums[i]%2 != nums[i-1]%2)
	}

	// Step 2: Precompute prefix sums for fast range queries
	prefixSum := make([]int, n)
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1]
		if !isSpecialPrefix[i] {
			prefixSum[i]++
		}
	}

	// Step 3: Process each query
	result := make([]bool, len(queries))
	for i, query := range queries {
		from, to := query[0], query[1]
		if from == to { // Single-element subarray is always special
			result[i] = true
		} else {
			// Check the prefix sum difference for the range
			specialCount := prefixSum[to] - prefixSum[from]
			result[i] = (specialCount == 0)
		}
	}

	return result
}
