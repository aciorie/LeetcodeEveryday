package codetop3rd

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	a, b := 1, 2
	for i := 3; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
