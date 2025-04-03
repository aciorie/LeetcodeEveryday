package hot2nd

import "reflect"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	freq := make(map[rune]int)
	for _, w := range t {
		freq[w]++
	}

	minLength, minWindowStr := -1, ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			subStr := s[i : j+1]

			// Check if the substring contains all the characters in t
			if contains(subStr, freq) {
				// If included, update the minimum window
				if minLength == -1 || len(subStr) < minLength {
					minLength = len(subStr)
					minWindowStr = subStr
				}
			}
		}
	}
	return minWindowStr
}

func contains(subStr string, freq map[rune]int) bool {
	freqSub := make(map[rune]int)
	for _, w := range subStr {
		freqSub[w]++
	}
	if !reflect.DeepEqual(freqSub, freq) {
		return false
	}
	return true
}
