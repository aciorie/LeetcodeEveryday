package codetop3rd

func searchRange(nums []int, target int) []int {
	leftBound := findLeftBound(nums, target)
	rightBound := findRightBound(nums, target)

	return []int{leftBound, rightBound}
}

func findLeftBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	res := -1
	for l <= r {
		m := l + (r-l)/2
		if target == nums[m] {
			res = m
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return res
}

func findRightBound(nums []int, target int) int {
	l, r := 0, len(nums)-1
	res := -1
	for l <= r {
		m := l + (r-l)/2
		if target == nums[m] {
			res = m
			l = m + 1
		} else if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return res
}
