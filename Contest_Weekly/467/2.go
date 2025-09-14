package Contest

import "sort"

func maxKDistinct(nums []int, k int) []int {
	sort.Ints(nums)
	n := len(nums)
	res := make([]int, 0, k)

	for i := n - 1; i >= 0; i-- {
		if len(res) == k {
			return res
		}

		if i < n-1 && nums[i] == nums[i+1] {
			continue
		}

		res = append(res, nums[i])
	}

	return res
}
