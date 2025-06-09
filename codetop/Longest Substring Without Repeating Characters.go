package codetop

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	res := 0
	windowStart := 0
	charLastIndex := make(map[rune]int)

	for windowEnd, currentChar := range s {
		if lastPos, found := charLastIndex[currentChar]; found && lastPos >= windowStart {
			windowStart = lastPos + 1
		}

		charLastIndex[currentChar] = windowEnd

		currentWindowLength := windowEnd - windowStart + 1
		if currentWindowLength > res {
			res = currentWindowLength
		}
	}

	return res
}
