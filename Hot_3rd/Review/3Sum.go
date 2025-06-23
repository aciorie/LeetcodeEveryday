package review

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	// i作为最左边的元素
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// nums[i]都为正数的话后面都是正数
		if nums[i] > 0 {
			break
		}

		m, r := i+1, len(nums)-1
		for m < r {
			curSum := nums[i] + nums[m] + nums[r]
			if curSum == 0 {
				res = append(res, []int{nums[i], nums[m], nums[r]})
				m++
				r--

				for m < r && nums[m] == nums[m-1] {
					m++
				}
				for m < r && nums[r] == nums[r+1] {
					r--
				}
			} else if curSum < 0 {
				m++
			} else {
				r--
			}
		}
	}
	return res
}
