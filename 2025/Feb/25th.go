package feb

/*
1524. Number of Sub-arrays With Odd Sum
Medium
Topics
Companies
Hint
Given an array of integers arr, return the number of subarrays with an odd sum.

Since the answer can be very large, return it modulo 109 + 7.



Example 1:

Input: arr = [1,3,5]
Output: 4
Explanation: All subarrays are [[1],[1,3],[1,3,5],[3],[3,5],[5]]
All sub-arrays sum are [1,4,9,3,8,5].
Odd sums are [1,9,3,5] so the answer is 4.
Example 2:

Input: arr = [2,4,6]
Output: 0
Explanation: All subarrays are [[2],[2,4],[2,4,6],[4],[4,6],[6]]
All sub-arrays sum are [2,6,12,4,10,6].
All sub-arrays have even sum and the answer is 0.
Example 3:

Input: arr = [1,2,3,4,5,6,7]
Output: 16


Constraints:

1 <= arr.length <= 105
1 <= arr[i] <= 100
*/

const MOD = 1_000_000_007

//failed
func NumOfSubarrays(arr []int) int {
	oddNum := 0
	for _, v := range arr {
		if v%2 == 1 {
			oddNum++
		}
	}
	res := oddNum

	for i := 1; i < oddNum; i += 2 {
		// dont know how to continue
	}
	return res
}

func NumOfSubarrays2(arr []int) int {
	odd, even, curSum, res := 0, 1, 0, 0
	for _, v := range arr {
		curSum += v
		if curSum%2 == 0 {
			res = (res + odd) % MOD
			even++
		} else {
			res = (res + even) % MOD
			odd++
		}
	}
	return res
}
