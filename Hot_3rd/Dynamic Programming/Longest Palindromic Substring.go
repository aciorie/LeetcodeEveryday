package main

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	start, maxLength := 0, 1
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for i := 0; i < n; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if length > maxLength {
					maxLength = length
					start = i
				}
			}
		}
	}

	return s[start : start+maxLength]
}
