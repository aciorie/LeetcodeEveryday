package review

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
		if nums[m] == target {
			return m
		}
		if nums[l] < nums[m] {
			if nums[m] < target && target < nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if nums[l] < target && target < nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
