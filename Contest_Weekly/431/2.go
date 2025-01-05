package contest

/*
Q2. Find Mirror Score of a String
Medium
4 pt.
You are given a string s.

We define the mirror of a letter in the English alphabet as its corresponding letter when the alphabet is reversed. For example, the mirror of 'a' is 'z', and the mirror of 'y' is 'b'.

Initially, all characters in the string s are unmarked.

You start with a score of 0, and you perform the following process on the string s:

Iterate through the string from left to right.
At each index i, find the closest unmarked index j such that j < i and s[j] is the mirror of s[i]. Then, mark both indices i and j, and add the value i - j to the total score.
If no such index j exists for the index i, move on to the next index without making any changes.
Return the total score at the end of the process.



Example 1:
Input: s = "aczzx"
Output: 5
Explanation:
i = 0. There is no index j that satisfies the conditions, so we skip.
i = 1. There is no index j that satisfies the conditions, so we skip.
i = 2. The closest index j that satisfies the conditions is j = 0, so we mark both indices 0 and 2, and then add 2 - 0 = 2 to the score.
i = 3. There is no index j that satisfies the conditions, so we skip.
i = 4. The closest index j that satisfies the conditions is j = 1, so we mark both indices 1 and 4, and then add 4 - 1 = 3 to the score.


Example 2:
Input: s = "abcdef"
Output: 0
Explanation:
For each index i, there is no index j that satisfies the conditions.



Constraints:

1 <= s.length <= 105
s consists only of lowercase English letters.
*/

/*
Wrong Answer	404 / 581 testcases passed
Input	s =	"azapfwonwwcdagew"
Output	7		Expected	3
*/
func CalculateScore(s string) int64 {
	if len(s) == 1 {
		return int64(0)
	}

	var res int64
	marked := make([]bool, len(s))

	for i := 0; i < len(s); i++ {
		if marked[i] {
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if !marked[j] && isMirror(s[i], s[j]) {
				res += int64(j - i)
				marked[i] = true
				marked[j] = true
				break
			}
		}
	}
	return res
}

func isMirror(letter1, letter2 byte) bool {
	return letter1+letter2 == 'a'+'z'
}
