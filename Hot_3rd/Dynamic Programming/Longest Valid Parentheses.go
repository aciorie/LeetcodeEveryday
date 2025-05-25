package main

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	stack := []int{-1}
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				currentLen := i - stack[len(stack)-1]
				if currentLen > maxLen {
					maxLen = currentLen
				}
			}
		}
	}
	return maxLen
}
