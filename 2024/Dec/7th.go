package Dec

import "math"

/*
Minimum Limit of Balls in a Bag
Medium
Topics
Companies
Hint
You are given an integer array nums where the ith bag contains nums[i] balls. You are also given an integer maxOperations.

You can perform the following operation at most maxOperations times:

Take any bag of balls and divide it into two new bags with a positive number of balls.
For example, a bag of 5 balls can become two new bags of 1 and 4 balls, or two new bags of 2 and 3 balls.
Your penalty is the maximum number of balls in a bag. You want to minimize your penalty after the operations.

Return the minimum possible penalty after performing the operations.



Example 1:

Input: nums = [9], maxOperations = 2
Output: 3
Explanation:
- Divide the bag with 9 balls into two bags of sizes 6 and 3. [9] -> [6,3].
- Divide the bag with 6 balls into two bags of sizes 3 and 3. [6,3] -> [3,3,3].
The bag with the most number of balls has 3 balls, so your penalty is 3 and you should return 3.
Example 2:

Input: nums = [2,4,8,2], maxOperations = 4
Output: 2
Explanation:
- Divide the bag with 8 balls into two bags of sizes 4 and 4. [2,4,8,2] -> [2,4,4,4,2].
- Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,4,4,4,2] -> [2,2,2,4,4,2].
- Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,2,2,4,4,2] -> [2,2,2,2,2,4,2].
- Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,2,2,2,2,4,2] -> [2,2,2,2,2,2,2,2].
The bag with the most number of balls has 2 balls, so your penalty is 2, and you should return 2.


Constraints:

1 <= nums.length <= 105
1 <= maxOperations, nums[i] <= 109
*/

func MinimumSize(nums []int, maxOperations int) int {
	l, r := 0, maxInArray(nums)
	res := r
	for l <= r {
		//m represents the maximum bag size we try
		m := l + (r-l)/2
		//check if the maximum size of all bags can be reduced to
		//m in at most maxOperations operations.
		if canAchieve(nums, m, maxOperations) {
			res = m
			//try a smaller value
			r = m - 1
		} else {
			//try a bigger value
			l = m + 1
		}
	}
	return res
}

func maxInArray(nums []int) int {
	ma := math.MinInt32
	for _, num := range nums {
		if ma < num {
			ma = num
		}
	}
	return ma
}

//check whether a bag of size maxSize can be achieved in at most maxOperations operations
func canAchieve(nums []int, maxSize, maxOperations int) bool {
	if maxSize <= 0 {
		return false
	}

	times := 0
	for _, num := range nums {
		if num > maxSize {
			//calculates the number of operations required to split num to maxSize
			times += (num - 1) / maxSize
		}
	}
	return times <= maxOperations
}
