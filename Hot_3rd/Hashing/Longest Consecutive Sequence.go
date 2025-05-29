package hashing

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	numSet, res := make(map[int]bool), 0
	for _, num := range nums {
		numSet[num] = true
	}
	for _, num := range nums {
		if !numSet[num-1] {
			curNum := num
			curLen := 1
			for numSet[curNum+1] {
				curNum++
				curLen++
			}
			if res < curLen {
				res = curLen
			}
		}
	}
	return res
}
