package feb

/*
1749. Maximum Absolute Sum of Any Subarray
Medium
Topics
Companies
Hint
You are given an integer array nums. The absolute sum of a subarray [numsl, numsl+1, ..., numsr-1, numsr] is abs(numsl + numsl+1 + ... + numsr-1 + numsr).

Return the maximum absolute sum of any (possibly empty) subarray of nums.

Note that abs(x) is defined as follows:

If x is a negative integer, then abs(x) = -x.
If x is a non-negative integer, then abs(x) = x.


Example 1:

Input: nums = [1,-3,2,3,-4]
Output: 5
Explanation: The subarray [2,3] has absolute sum = abs(2+3) = abs(5) = 5.
Example 2:

Input: nums = [2,-5,1,-4,3,-2]
Output: 8
Explanation: The subarray [-5,1,-4] has absolute sum = abs(-5+1-4) = abs(-8) = 8.


Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
*/
func MaxAbsoluteSum(nums []int) int {
	if len(nums) == 1 {
		return abs(nums[0])
	}
	maxSoFar := 0
	minSoFar := 0
	maxAbsSum := 0

	for _, num := range nums {
		maxSoFar = max(num, maxSoFar+num)
		minSoFar = min(num, minSoFar+num)
		currentAbsMax := max(abs(maxSoFar), abs(minSoFar))
		if currentAbsMax > maxAbsSum {
			maxAbsSum = currentAbsMax
		}
	}
	return maxAbsSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
