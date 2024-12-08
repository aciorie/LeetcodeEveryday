package hot

/*
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.



Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]


Constraints:

n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.


Follow up: Could you come up with a one-pass algorithm using only constant extra space?
*/

//two pointers move towards the center from the beginning and the end of the original array
func SortColors(nums []int) {
	zeros, ones, twos := 0, 0, len(nums)-1

	for ones <= twos {
		if nums[ones] == 0 {
			nums[zeros], nums[ones] = nums[ones], nums[zeros]
			zeros++
			ones++
		} else if nums[ones] == 1 {
			ones++
		} else {
			nums[ones], nums[twos] = nums[twos], nums[ones]
			twos--
		}
	}
}
