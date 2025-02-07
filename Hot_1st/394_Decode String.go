package hot

import (
	"strconv"
	"strings"
)

/*
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 105.



Example 1:

Input: s = "3[a]2[bc]"
Output: "aaabcbc"
Example 2:

Input: s = "3[a2[c]]"
Output: "accaccacc"
Example 3:

Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"


Constraints:

1 <= s.length <= 30
s consists of lowercase English letters, digits, and square brackets '[]'.
s is guaranteed to be a valid input.
All the integers in s are in the range [1, 300].
*/

// failed
func DecodeString1(s string) string {
	var res strings.Builder
	for i := 0; i < len(s)-1; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			res.WriteByte(s[i])
		}
		if s[i] >= '1' && s[i] <= '9' {
			for j := 0; j <= int(s[i]); j++ {
				//dont know how to continue
			}
		}
	}

	return res.String()
}

func DecodeString2(s string) string {
	var res, numStr strings.Builder
	var stack []string

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			res.WriteByte(s[i])
		} else if s[i] >= '1' && s[i] <= '9' {
			numStr.WriteByte(s[i])
		} else if s[i] == '[' {
			stack = append(stack, res.String())
			stack = append(stack, numStr.String())

			res.Reset()
			numStr.Reset()
		} else if s[i] == ']' {
			countStr := stack[len(stack)-1] // Remove the number of repetitions at the top of the stack
			stack = stack[:len(stack)-1]
			prevStr := stack[len(stack)-1] // Remove the previously decoded string at the top of the stack
			stack = stack[:len(stack)-1]

			count, _ := strconv.Atoi(countStr)

			for j := 0; j < count; j++ {
				prevStr += res.String()
			}
			res.Reset()
			res.WriteString(prevStr)
		}
	}

	return res.String()
}
