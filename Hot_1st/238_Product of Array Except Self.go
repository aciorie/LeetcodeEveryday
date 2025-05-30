package hot

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]


Constraints:
2 <= nums.length <= 105
-30 <= nums[i] <= 30
The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.


Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	forward, backward := make([]int, n), make([]int, n)
	res := make([]int, n)

	// this would make the result 4 times bigger than expected
	// forward[0], backward[n-1] = nums[0], nums[n-1]

	forward[0], backward[n-1] = 1, 1

	for i := 0; i < n-1; i++ {
		forward[i+1] = forward[i] * nums[i]
	}
	for i := n - 1; i > 0; i-- {
		backward[i-1] = backward[i] * nums[i]
	}

	for i := 0; i < n; i++ {
		res[i] = forward[i] * backward[i]
	}
	return res
}
