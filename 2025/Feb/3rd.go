package feb

import "sort"

/*
3105. Longest Strictly Increasing or Strictly Decreasing Subarray
Easy
Topics
Companies
You are given an array of integers nums. Return the length of the longest
subarray
 of nums which is either
strictly increasing
 or
strictly decreasing
.



Example 1:

Input: nums = [1,4,3,3,2]

Output: 2

Explanation:

The strictly increasing subarrays of nums are [1], [2], [3], [3], [4], and [1,4].

The strictly decreasing subarrays of nums are [1], [2], [3], [3], [4], [3,2], and [4,3].

Hence, we return 2.

Example 2:

Input: nums = [3,3,3,3]

Output: 1

Explanation:

The strictly increasing subarrays of nums are [3], [3], [3], and [3].

The strictly decreasing subarrays of nums are [3], [3], [3], and [3].

Hence, we return 1.

Example 3:

Input: nums = [3,2,1]

Output: 3

Explanation:

The strictly increasing subarrays of nums are [3], [2], and [1].

The strictly decreasing subarrays of nums are [3], [2], [1], [3,2], [2,1], and [3,2,1].

Hence, we return 3.



Constraints:

1 <= nums.length <= 50
1 <= nums[i] <= 50

*/

/*
Wrong Answer
251 / 868 testcases passed

Editorial
Input
nums =
[1,4,3,3,2]

Use Testcase
Output
3
Expected
2
*/
func LongestMonotonicSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	res := 1
	sort.Ints(nums)
	curMax := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i]+1 {
			curMax++
		} else {
			curMax = 1
		}
		res = max(res, curMax)
	}
	return res
}

func LongestMonotonicSubarray2(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	res := 1
	curMax := 1
	curMin := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curMax++
			curMin = 1
		} else if nums[i] < nums[i-1] {
			curMin++
			curMax = 1
		} else {
			curMin = 1
			curMax = 1
		}
		res = max(res, max(curMax, curMin))
	}
	return res
}
