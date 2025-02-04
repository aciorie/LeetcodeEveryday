package feb

/*
1800. Maximum Ascending Subarray Sum
Easy
Topics
Companies
Hint
Given an array of positive integers nums, return the maximum possible sum of an ascending subarray in nums.

A subarray is defined as a contiguous sequence of numbers in an array.

A subarray [numsl, numsl+1, ..., numsr-1, numsr] is ascending if for all i where l <= i < r, numsi  < numsi+1. Note that a subarray of size 1 is ascending.



Example 1:

Input: nums = [10,20,30,5,10,50]
Output: 65
Explanation: [5,10,50] is the ascending subarray with the maximum sum of 65.
Example 2:

Input: nums = [10,20,30,40,50]
Output: 150
Explanation: [10,20,30,40,50] is the ascending subarray with the maximum sum of 150.
Example 3:

Input: nums = [12,17,15,13,10,11,12]
Output: 33
Explanation: [10,11,12] is the ascending subarray with the maximum sum of 33.


Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/

/*
Wrong Answer
50 / 104 testcases passed

Editorial
Input
nums =
[100,10,1]

Use Testcase
Output
10
Expected
100
*/
func MaxAscendingSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res, n := 0, len(nums)-1
	curMax := nums[0]
	for i := 0; i < n; i++ {
		if nums[i] > nums[i+1] {
			curMax = nums[i+1]
		} else {
			curMax += nums[i+1]
		}
		res = max(res, curMax)
	}
	return res
}

/*
Wrong Answer
57 / 104 testcases passed

Editorial
Input
nums =
[3,6,10,1,8,9,9,8,9]

Use Testcase
Output
27
Expected
19
*/
func MaxAscendingSum2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res, n := nums[0], len(nums)-1
	curMax := nums[0]
	for i := 0; i < n; i++ {
		if nums[i] > nums[i+1] {
			curMax = nums[i+1]
		} else {
			curMax += nums[i+1]
		}
		res = max(res, curMax)
	}
	return res
}

func MaxAscendingSum3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res, curMax := nums[0], nums[0]

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			curMax += nums[i+1]
		} else {
			curMax = nums[i+1]
		}
		res = max(res, curMax)
	}
	return res
}
