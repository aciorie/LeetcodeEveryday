package hot2nd

func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		m := l + (r-l)/2
		if target == nums[m] {
			return m
		} else if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}
