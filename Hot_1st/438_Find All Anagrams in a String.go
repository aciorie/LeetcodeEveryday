package hot

import "reflect"

/*
Given two strings s and p, return an array of all the start indices of p's
anagrams
 in s. You may return the answer in any order.



Example 1:
Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".

Example 2:
Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".


Constraints:

1 <= s.length, p.length <= 3 * 104
s and p consist of lowercase English letters.

*/
func FindAnagrams(s string, p string) []int {
	var res []int
	if len(p) > len(s) {
		return res
	}

	occur, n := make(map[rune]int), len(p)
	for _, cha := range p {
		occur[cha]++
	}

	occurInS := make(map[byte]int, n)
	for i := 0; i < n; i++ {
		occurInS[s[i]]++
	}

	i := n
	for j := n + n; j < len(s); {
		if reflect.DeepEqual(occur, occurInS) {
			res = append(res, i)
		}

		//pop the first element
		// occurInS=[1:]
		i++
		j++
		//push the last element
		// occoccurInS
	}

	return res
}

func FindAnagrams2(s string, p string) []int {
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
