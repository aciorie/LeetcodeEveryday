package codetop2nd

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		m, r := i+1, n-1
		for m < r {
			cursum := nums[i] + nums[m] + nums[r]
			if cursum == 0 {
				res = append(res, []int{nums[i], nums[m], nums[r]})
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
				m++
			}
		}

	}

	return res
}
