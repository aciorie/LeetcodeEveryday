package main

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	n := len(nums)
	l, r := 0, n-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[r] {
			l = m + 1
		} else if nums[m] < nums[r] {
			r = m - 1
		} else {
			r--
		}
	}
	return nums[l]
}
