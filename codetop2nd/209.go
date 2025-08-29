package codetop2nd

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := n + 1
	left, sum := 0, 0

	for right, num := range nums {
		sum += num
		for sum >= target {
			curLen := right - left + 1
			res = min(res, curLen)
		}

		sum -= nums[left]
		left++
	}

	if res == n+1 {
		return 0
	}
	return res
}
