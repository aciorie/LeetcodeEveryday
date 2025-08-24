package codetop2nd

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	a, b := 1, 2
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
