package boring

import "sort"

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sort.Ints(nums)
	leftSum, rightSum, l, r := nums[0], nums[n-1], 1, n-2
	for l <= r {
		if leftSum == rightSum {
			return true
		} else if leftSum < rightSum {
			leftSum += nums[l]
			l++
		} else {
			rightSum += nums[r]
			r--
		}
	}
	return false
}
