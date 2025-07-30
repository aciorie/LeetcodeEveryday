package hashing

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numSet, res := make(map[int]bool), 0
	for _, num := range nums {
		numSet[num] = true
	}

	for _, num := range nums {
		if !numSet[num] {
			curNum, curLen := num, 1
			for numSet[curNum+1] {
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
