package leetcode75

func pivotIndex(nums []int) int {
	n := len(nums)

	leftSum, totalSum := 0, 0
	for _, num := range nums {
		totalSum += num
	}

	for i := 0; i < n; i++ {
		rightSum := totalSum - leftSum - nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
