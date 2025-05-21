package hot3rd

// false one
func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	n := len(nums)
	l, r := 0, n-1

	for l <= r {
		m := l + (r-l)/2
		if m == target {
			return m
		}

		if nums[l] <= nums[m] {
			// 左边有序看右边
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			// 右边有序看左边
			if nums[l] < target && target <= nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}
