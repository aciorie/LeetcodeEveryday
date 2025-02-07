package hot

/*
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.



Example 1:
Input: nums = [1,1,1], k = 2
Output: 2

Example 2:
Input: nums = [1,2,3], k = 3
Output: 2


Constraints:

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/
func SubarraySum(nums []int, k int) int {
	prefix, res := 0, 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == k && nums[i+1] > nums[i] {
			res++
			break
		}
		if nums[i] == k && nums[i+1] == nums[i] {
			res++
		}
		if nums[i] < k {
			prefix += nums[i]
			if prefix == k {
				res++
			}
		}
	}

	return res
}
