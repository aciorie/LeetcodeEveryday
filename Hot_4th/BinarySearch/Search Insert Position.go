package binarysearch

func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		if nums[0] < target {
			return 1
		} else {
			return 0
		}
	}

	l, r := 0, n-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
