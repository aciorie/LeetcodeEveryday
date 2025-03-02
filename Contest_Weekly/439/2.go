package contest

import "strconv"

/*
Q2. Longest Palindromic Subsequence After at Most K Operations
Medium
4 pt.
You are given a string s and an integer k.

In one operation, you can replace the character at any position with the next or previous letter in the alphabet (wrapping around so that 'a' is after 'z'). For example, replacing 'a' with the next letter results in 'b', and replacing 'a' with the previous letter results in 'z'. Similarly, replacing 'z' with the next letter results in 'a', and replacing 'z' with the previous letter results in 'y'.

Return the length of the longest palindromic subsequence of s that can be obtained after performing at most k operations.

A subsequence is a non-empty string that can be derived from another string by deleting some or no characters without changing the order of the remaining characters.

A palindrome is a string that reads the same forward and backward.



Example 1:

Input: s = "abced", k = 2

Output: 3

Explanation:

Replace s[1] with the next letter, and s becomes "acced".
Replace s[4] with the previous letter, and s becomes "accec".
The subsequence "ccc" forms a palindrome of length 3, which is the maximum.

Example 2:

Input: s = "aaazzz", k = 4

Output: 6

Explanation:

Replace s[0] with the previous letter, and s becomes "zaazzz".
Replace s[4] with the next letter, and s becomes "zaazaz".
Replace s[3] with the next letter, and s becomes "zaaaaz".
The entire string forms a palindrome of length 6.



Constraints:

1 <= s.length <= 200
1 <= k <= 200
s consists of only lowercase English letters.
*/

//TLE
func LongestPalindromicSubsequence(s string, k int) int {
	memo := make(map[string]int)
	return dpRec(s, 0, len(s)-1, k, memo)
}

// dpRec uses memoized search to find the value of dp(i, j, remain),
// memo is used for caching, the key format is "i_j_remain"
func dpRec(s string, i, j, remain int, memo map[string]int) int {
	// Construction cache key
	key := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(remain)
	if val, exists := memo[key]; exists {
		return val
	}

	// Boundary conditions: No characters
	if i > j {
		return 0
	}
	// Single character, total length is 1, no operation required
	if i == j {
		return 1
	}

	best := 0

	// Calculate the operation cost of s[i] and s[j]:
	// cost = min(|s[i]-s[j]|, 26 - |s[i]-s[j]|)
	diff := abs(int(s[i]) - int(s[j]))
	cost := min(diff, 26-diff)

	// If an operation can be used to turn s[i] and s[j] into the same character, then it is considered a match.
	if cost <= remain {
		// After matching, each side contributes 1, and the contribution in the interval is dp(i+1, j-1, remain-cost)
		best = 2 + dpRec(s, i+1, j-1, remain-cost, memo)
	}

	// Skip characters on the left
	best = max(best, dpRec(s, i+1, j, remain, memo))
	// Skip right character
	best = max(best, dpRec(s, i, j-1, remain, memo))

	// Save the results to the cache
	memo[key] = best
	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
