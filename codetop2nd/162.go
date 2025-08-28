package codetop2nd

func findPeakElement(nums []int) int {
	n := len(nums)
	if n < 1 {
		return -1
	}
	l, r := 0, n-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
