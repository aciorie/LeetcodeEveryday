package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// 优化：如果第一个元素已经大于 0，那么后面的元素也都会大于 0
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 去重
		if nums[i] > 0 {
			break
		}

		l, r := i+1, len(nums)-1
		for l < r {
			curSum := nums[i] + nums[l] + nums[r]
			if curSum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--

				// 去重
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if curSum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	
	return res
}
