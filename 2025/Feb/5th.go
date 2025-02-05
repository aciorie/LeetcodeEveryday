package feb

/*
1790. Check if One String Swap Can Make Strings Equal
Easy
Topics
Companies
Hint
You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.

Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.



Example 1:

Input: s1 = "bank", s2 = "kanb"
Output: true
Explanation: For example, swap the first character with the last character of s2 to make "bank".
Example 2:

Input: s1 = "attack", s2 = "defend"
Output: false
Explanation: It is impossible to make them equal with one string swap.
Example 3:

Input: s1 = "kelb", s2 = "kelb"
Output: true
Explanation: The two strings are already equal, so no string swap operation is required.


Constraints:

1 <= s1.length, s2.length <= 100
s1.length == s2.length
s1 and s2 consist of only lowercase English letters.
*/

/*
Wrong Answer
72 / 132 testcases passed

Editorial
Input
s1 =
"bank"
s2 =
"kanb"

Use Testcase
Output
false
Expected
true
*/
func AreAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	stack := make([]byte, 2)
	times := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		if s1[i] != s2[i] {
			if times > 1 {
				return false
			}
			if len(stack) == 0 {
				stack = append(stack, s1[i])
				stack = append(stack, s2[i])
			} else {
				if s2[i] == stack[0] && s1[i] == stack[1] {
					stack = stack[:0]
					times++
				} else {
					return false
				}
			}
		}
	}
	return true
}

func AreAlmostEqual2(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	var diff []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}

	if len(diff) == 2 {
		return s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]]
	}

	return false
}
