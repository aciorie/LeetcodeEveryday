package contest

import "sort"

/*
Q2. Maximum Number of Distinct Elements After Operations		Medium

You are given an integer array nums and an integer k.

You are allowed to perform the following operation on each element of the array at most once:

Add an integer in the range [-k, k] to the element.
Return the maximum possible number of distinct elements in nums after performing the operations.



Example 1:
Input: nums = [1,2,2,3,3,4], k = 2
Output: 6
Explanation:
nums changes to [-1, 0, 1, 2, 3, 4] after performing operations on the first four elements.

Example 2:
Input: nums = [4,4,4,4], k = 1
Output: 3
Explanation:
By adding -1 to nums[0] and 1 to nums[1], nums changes to [3, 5, 4, 4].



Constraints:

1 <= nums.length <= 105
1 <= nums[i] <= 109
0 <= k <= 109
*/

//Wrong Answer		579 / 630 testcases passed
func MaxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	distinctSet := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := distinctSet[num]; !exists {
			distinctSet[num] = struct{}{}
		} else {
			for i := 1; i <= k; i++ {
				if _, exists := distinctSet[num-i]; !exists {
					distinctSet[num-i] = struct{}{}
					break
				}
				if _, exists := distinctSet[num+i]; !exists {
					distinctSet[num+i] = struct{}{}
					break
				}
			}
		}
	}

	return len(distinctSet)
}