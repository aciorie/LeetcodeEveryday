package jul

func maxSum(nums []int) int {
	sum, hasPositive, seen := 0, false, make(map[int]bool)

	for _, num := range nums {
		if num > 0 {
			if !seen[num] {
				sum += num
				hasPositive = true
				seen[num] = true
			}
		}
	}

	if hasPositive {
		return sum
	}

	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}
