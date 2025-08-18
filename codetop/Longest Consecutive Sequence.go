package codetop

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	seen := make([]bool, n+1)
	for _, num := range nums {
		seen[num] = true
	}

	res := 0
	for _, num := range nums {
		if !seen[num-1] {
			// 说明 num 是一个连续序列的起点
			curNum, curLen := num, 1
			for seen[curNum+1] {
				curLen++
				curNum++
			}
			if curLen > res {
				res = curLen
			}
		}
	}
	return res
}
