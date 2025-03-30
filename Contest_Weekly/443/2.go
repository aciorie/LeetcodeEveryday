package Contest

/*
Q2. Longest Palindrome After Substring Concatenation I
Medium
4 pt.
You are given two strings, s and t.

You can create a new string by selecting a substring from s (possibly empty) and a substring from t (possibly empty), then concatenating them in order.

Return the length of the longest palindrome that can be formed this way.

A substring is a contiguous sequence of characters within a string.

A palindrome is a string that reads the same forward and backward.



Example 1:

Input: s = "a", t = "a"

Output: 2

Explanation:

Concatenating "a" from s and "a" from t results in "aa", which is a palindrome of length 2.

Example 2:

Input: s = "abc", t = "def"

Output: 1

Explanation:

Since all characters are different, the longest palindrome is any single character, so the answer is 1.

Example 3:

Input: s = "b", t = "aaaa"

Output: 4

Explanation:

Selecting "aaaa" from t is the longest palindrome, so the answer is 4.

Example 4:

Input: s = "abcde", t = "ecdba"

Output: 5

Explanation:

Concatenating "abc" from s and "ba" from t results in "abcba", which is a palindrome of length 5.



Constraints:

1 <= s.length, t.length <= 30
s and t consist of lowercase English letters.
*/
func LongestPalindrome(s string, t string) int {
	res := 0

	// Generate all substrings of s
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			substringS := s[i:j]

			// Generate all substrings of t
			for k := 0; k <= len(t); k++ {
				for l := k; l <= len(t); l++ {
					substringT := t[k:l]

					// Concatenate substrings
					combined := substringS + substringT

					if isPalindrome(combined) {
						if len(combined) > res {
							res = len(combined)
						}
					}
				}
			}
		}
	}

	return res
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
