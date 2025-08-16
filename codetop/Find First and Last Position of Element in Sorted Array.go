package codetop

func searchRange(nums []int, target int) []int {
	// n := len(nums)
	// res := []int{-1, -1}

	// l, r := 0, n-1
	// for l < r {
	// 	m := l + (r-l)/2
	// 	if target == nums[m] {
	// 		tmp := nums[m]
	// 		for tmp > 0 {
	// 			if nums[tmp] == nums[m] {
	// 				tmp--
	// 			}
	// 		}
	// 		for m < r {
	// 			if nums[m] == nums[m-1] {
	// 				m++
	// 			}
	// 		}
	// 		res[0], res[1] = tmp, m
	// 	} else if nums[m] < target {
	// 		l = m + 1
	// 	} else {
	// 		r = m - 1
	// 	}
	// }
	// return res
	leftmost := findBound(nums, target, true)
	if leftmost == -1 {
		return []int{-1, -1}
	}
	rightmost := findBound(nums, target, false)

	return []int{leftmost, rightmost}
}

func findBound(nums []int, target int, isLeftBound bool) int {
	low, high, res := 0, len(nums)-1, -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			res = mid
			if isLeftBound {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}
