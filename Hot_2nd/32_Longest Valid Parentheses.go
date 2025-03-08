package hot2nd

// failed one
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	n := len(s)
	dp, res := make([]int, n), 0

	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 { // 考虑越界情况，例如 "()"
					dp[i] = dp[i-2] + 2
				} else { // i < 2, 例如 "()"，此时 i-2 = -1 或 -2, dp[i-2] 会越界
					dp[i] = 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {

			}
		}
	}
	return res
}
