package codetop2nd

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	l, r := 1, x
	res := 1
	for l <= r {
		m := l + (r-l)/2
		// 避免 mid*mid 导致的整数溢出
		if m > x/m {
			r = m - 1
		} else {
			res = m
			l = m + 1
		}
	}
	return res
}
