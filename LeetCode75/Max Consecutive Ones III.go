package leetcode75

func longestOnes(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		if k > 0 {
			return 1
		} else {
			return 0
		}
	}
	l, r, zeroCount, maxLength := 0, 0, 0, 0
	for r < n {
		if nums[r] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[l] == 0 {
				zeroCount--
			}
			l++
		}

		maxLength = max(maxLength, r-l+1)
		r++
	}
	return maxLength
}
