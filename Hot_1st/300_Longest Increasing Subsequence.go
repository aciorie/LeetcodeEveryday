package hot

/*
Given an integer array nums, return the length of the longest strictly increasing
subsequence
.



Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1


Constraints:
1 <= nums.length <= 2500
-104 <= nums[i] <= 104


Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?
*/

/*
Wrong Answer
35 / 55 testcases passed

Editorial
Input
nums =
[1,3,6,7,9,4,10,5,6]

Use Testcase
Output
5
Expected
6
*/
func LengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	//Each element itself can form a subsequence of length 1
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(dp)-1]
}

//	want:5,got:6
func LengthOfLIS2(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	//Each element itself can form a subsequence of length 1
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for _, length := range dp {
		if length > res {
			res = length
		}
	}

	return res
}
