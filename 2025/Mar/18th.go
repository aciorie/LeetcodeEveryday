package mar

/*
2401. Longest Nice Subarray
Medium
Topics
Companies
Hint
You are given an array nums consisting of positive integers.

We call a subarray of nums nice if the bitwise AND of every pair of elements that are in different positions in the subarray is equal to 0.

Return the length of the longest nice subarray.

A subarray is a contiguous part of an array.

Note that subarrays of length 1 are always considered nice.



Example 1:

Input: nums = [1,3,8,48,10]
Output: 3
Explanation: The longest nice subarray is [3,8,48]. This subarray satisfies the conditions:
- 3 AND 8 = 0.
- 3 AND 48 = 0.
- 8 AND 48 = 0.
It can be proven that no longer nice subarray can be obtained, so we return 3.
Example 2:

Input: nums = [3,1,5,11,13]
Output: 1
Explanation: The length of the longest nice subarray is 1. Any subarray of length 1 can be chosen.


Constraints:

1 <= nums.length <= 105
1 <= nums[i] <= 109
*/

// Wrong Answer		33 / 65 testcases passed
func LongestNiceSubarray(nums []int) int {
	l, res := 0, 1

	for r := 1; r < len(nums); r++ {
		currentAnd := nums[r]

		for i := r - 1; i >= l; i-- {
			if (currentAnd & nums[i]) != 0 {
				l = i + 1 // Move the left boundary to the element *after* the conflicting element
				break     // Exit the inner loop
			}
			currentAnd &= nums[i] // Update currentAnd
		}

		res = max(res, r-l)
	}
	return res
}
