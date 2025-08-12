package slidingwindow

import "reflect"

func findAnagrams(s string, p string) []int {
	var res []int
	if len(p) > len(s) {
		return res
	}

	occur := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		occur[p[i]]++
	}

	occurInS := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		occurInS[s[i]]++
	}

	for i := 0; i <= len(s)-len(p); i++ {
		if i > 0 {
			// Update occurInS
			occurInS[s[i-1]]-- // Remove the character to the left
			if occurInS[s[i-1]] == 0 {
				delete(occurInS, s[i-1]) // If the count is 0, delete the character
			}
			occurInS[s[i+len(p)-1]]++ // Add new characters on the right
		}

		// Checks if the current window is an anagram of p
		if reflect.DeepEqual(occur, occurInS) {
			res = append(res, i)
		}
	}
	return res
}
