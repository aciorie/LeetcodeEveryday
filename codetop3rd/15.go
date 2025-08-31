package codetop3rd

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for l := 0; l < len(nums); l++ {
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}

		if nums[l] > 0 {
			break
		}

		m, r := l+1, len(nums)-1
		for m < r {
			cursum := nums[l] + nums[m] + nums[r]
			if cursum == 0 {
				res = append(res, []int{nums[l], nums[m], nums[r]})
				m++
				r--

				for m < r && nums[m] == nums[m-1] {
					m++
				}
				for m < r && nums[r] == nums[r+1] {
					r--
				}
			} else if cursum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
