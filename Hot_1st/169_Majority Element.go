package hot

/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.



Example 1:
Input: nums = [3,2,3]
Output: 3

Example 2:
Input: nums = [2,2,1,1,1,2,2]
Output: 2


Constraints:
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109


Follow-up: Could you solve the problem in linear time and in O(1) space?
*/

//didn't solve in linear time and in O(1) space but got passed
func MajorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	res := nums[0]
	majorityCount := len(nums) / 2
	for k, v := range freq {
		if v > majorityCount {
			res = k
			break
		}
	}

	return res
}
