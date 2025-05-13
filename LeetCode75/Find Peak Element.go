package leetcode75

func findPeakElement(nums []int) int {
	n := len(nums)
	if n < 1 {
		return -1
	}
	resV, pos := nums[0], 0
	for i := 1; i < n; i++ {
		if nums[i] > resV {
			resV = nums[i]
			pos = i
		}
	}
	return pos
}

func findPeakElement2(nums []int) int {
	n := len(nums)
	if n < 1 {
		return -1
	}
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else if nums[mid] > nums[mid+1] {
			r = mid
		}
	}
	return l
}
