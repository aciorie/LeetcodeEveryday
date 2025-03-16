package contest

/*
Q1. Maximum Unique Subarray Sum After Deletion
Easy
3 pt.
You are given an integer array nums.

You are allowed to delete any number of elements from nums without making it empty. After performing the deletions, select a subarray of nums such that:

All elements in the subarray are unique.
The sum of the elements in the subarray is maximized.
Return the maximum sum of such a subarray.

A subarray is a contiguous non-empty sequence of elements within an array.


Example 1:

Input: nums = [1,2,3,4,5]

Output: 15

Explanation:

Select the entire array without deleting any element to obtain the maximum sum.

Example 2:

Input: nums = [1,1,0,1,1]

Output: 1

Explanation:

Delete the element nums[0] == 1, nums[1] == 1, nums[2] == 0, and nums[3] == 1. Select the entire array [1] to obtain the maximum sum.

Example 3:

Input: nums = [1,2,-1,-2,1,0,-1]

Output: 3

Explanation:

Delete the elements nums[2] == -1 and nums[3] == -2, and select the subarray [2, 1] from [1, 2, 1, 0, -1] to obtain the maximum sum.



Constraints:

1 <= nums.length <= 100
-100 <= nums[i] <= 100
*/

/*	Wrong Answer
191 / 927 testcases passed
Input
nums =
[-100]
Use Testcase
Output
0
Expected
-100
*/
// still failed
func MaxSum(nums []int) int {
	seen := make(map[int]bool)
	// l, cur, res := 0, 0, 0
	l, cur, res := 0, 0, nums[0]

	for r := 0; r < len(nums); r++ {
		// If the current element is already in the window
		for seen[nums[r]] {
			// Shrink the window from the left until the duplicate is removed
			cur -= nums[l]
			delete(seen, nums[l])
			l++
		}

		// Add the current element to the window
		seen[nums[r]] = true
		cur += nums[r]

		// Update maxSum if necessary
		if cur > res {
			res = cur
		}
	}
	return res
}

func MaxSum2(nums []int) int{
	sum := 0
	hasPositive := false
	// Use a map to record positive numbers we've added.
	seen := make(map[int]bool)
	
	for _, num := range nums {
		if num > 0 {
			// Add each positive number only once.
			if !seen[num] {
				sum += num
				seen[num] = true
				hasPositive = true
			}
		}
	}
	
	// If there is at least one positive, return the sum of unique positive numbers.
	if hasPositive {
		return sum
	}
	
	// Otherwise, all numbers are non-positive.
	// In that case, the best option is to choose the maximum element.
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}