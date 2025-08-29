package codetop2nd

func majorityElement(nums []int) int {
	n := len(nums)
	if n < 2 {
		return nums[0]
	}
	res, count := nums[0], 1
	for i := 1; i < n; i++ {
		if count == 0 {
			res = nums[i]
		}
		if nums[i] == res {
			count++
		} else {
			count--
		}
	}
	return res
}
