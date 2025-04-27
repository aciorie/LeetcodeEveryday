package leetcode75

import (
	"sort"
)

func maxOperations(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	l, r, res := 0, len(nums)-1, 0
	for l < r {
		if nums[l]+nums[r] == k {
			res++
			l++
			r--
		} else if nums[l]+nums[r] < k {
			l++
		} else {
			r--
		}
	}
	return res
}
