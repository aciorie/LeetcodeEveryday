package review

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	res, windowStart, charLastIndex := 0, 0, make(map[rune]int)

	for windowEnd, curChar := range s {
		if lastPos, exist := charLastIndex[curChar]; exist && lastPos >= windowStart {
			windowStart = lastPos + 1
		}

		charLastIndex[curChar] = windowEnd

		curWindowLength := windowEnd - windowStart + 1
		if curWindowLength > res {
			res = curWindowLength
		}
	}
	return res
}
