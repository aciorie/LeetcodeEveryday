package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	res, start := 0, 0
	ma := make(map[rune]int)

	for k, v := range s {
		if pos, ok := ma[v]; ok && pos >= start {
			start = pos + 1
		}
		ma[v] = k
		if k-start+1 > res {
			res = k - start + 1
		}
	}
	return res
}
