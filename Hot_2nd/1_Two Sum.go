package hot2nd

/*
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
*/
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	ma := make(map[int]int)
	for i, num := range nums {
		if index, exists := ma[target-num]; exists {
			res[0] = index
			res[1] = i
			return res
		}
		ma[num] = i
	}
	return res
}
