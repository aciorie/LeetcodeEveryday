package binarysearch

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	n := len(nums)

	if n == 0 {
		return res
	}

	l, r := 0, n-1
	firstIndex := -1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			firstIndex = m
			r = m - 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if firstIndex == -1 {
		return res
	}
	res[0] = firstIndex

	l, r = 0, n-1
	lastIndex := -1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			lastIndex = m
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	res[1] = lastIndex
	return res
}
