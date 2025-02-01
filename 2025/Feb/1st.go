package feb

/*
3151. Special Array I
Easy
Topics
Companies
Hint
An array is considered special if every pair of its adjacent elements contains two numbers with different parity.

You are given an array of integers nums. Return true if nums is a special array, otherwise, return false.



Example 1:
Input: nums = [1]
Output: true
Explanation:
There is only one element. So the answer is true.

Example 2:
Input: nums = [2,1,4]
Output: true
Explanation:
There is only two pairs: (2,1) and (1,4), and both of them contain numbers with different parity. So the answer is true.

Example 3:
Input: nums = [4,3,1,6]
Output: false
Explanation:
nums[1] and nums[2] are both odd. So the answer is false.



Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/

/*
Wrong Answer
803 / 852 testcases passed
Input	nums =	[1,5]
Output	true
Expected	false
*/
func IsArraySpecial(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums)-2; i++ {
		if (odd(nums[i]) && even(nums[i+1])) || (even(nums[i]) && odd(nums[i+1])) {
			continue
		} else {
			return false
		}
	}
	return true
}

/*
Wrong Answer
820 / 852 testcases passed
Input	nums =	[1,6,2]
Output	true
Expected	false
*/
func IsArraySpecial2(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	//added one statement
	if len(nums) == 2 {
		return (odd(nums[0]) && even(nums[1])) || (even(nums[0]) && odd(nums[1]))
	}
	for i := 0; i < len(nums)-2; i++ {
		if (odd(nums[i]) && even(nums[i+1])) || (even(nums[i]) && odd(nums[i+1])) {
			continue
		} else {
			return false
		}
	}
	return true
}

func IsArraySpecial3(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if len(nums) == 2 {
		return (odd(nums[0]) && even(nums[1])) || (even(nums[0]) && odd(nums[1]))
	}
	//changed here
	for i := 0; i < len(nums)-1; i++ {
		if (odd(nums[i]) && even(nums[i+1])) || (even(nums[i]) && odd(nums[i+1])) {
			continue
		} else {
			return false
		}
	}
	return true
}

func odd(i int) bool {
	return i%2 == 1
}

func even(i int) bool {
	return i%2 == 0
}
