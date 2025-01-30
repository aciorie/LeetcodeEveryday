package hot

import (
	"sort"
)

/*
Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or false otherwise.

Example 1:
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.

Constraints:

1 <= nums.length <= 200
1 <= nums[i] <= 100
*/

/*
Wrong Answer
73 / 144 testcases passed

Input	nums =[2,2,1,1]
Output	false
Expected	true
*/
func CanPartition(nums []int) bool {
	sort.Ints(nums)
	n, totalSum := len(nums), arraySum(nums)
	prevSum := 0

	for i := 0; i < n; i++ {
		prevSum += nums[i]
		totalSum -= nums[i]
		if prevSum == totalSum {
			return true
		}
	}
	return false
}

func arraySum(arr []int) int {
	total := 0
	for _, value := range arr {
		total += value
	}
	return total
}
