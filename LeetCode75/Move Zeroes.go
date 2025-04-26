package leetcode75

func moveZeroes(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}

	l, r := 0, 0
	for r < n-1 {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}
