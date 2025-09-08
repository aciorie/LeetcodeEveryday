package codetop3rd

func getNoZeroIntegers(n int) []int {
	var res []int
	var test_1317 func(a int) bool
	test_1317 = func(a int) bool {
		if a == 0 {
			return false
		}
		if a < 10 {
			return true
		}
		for a > 0 {
			if a%10 == 0 {
				return false
			}
			a /= 10
		}
		return true
	}
	
	for i := 1; i < n; i++ {
		if test_1317(i) != false && test_1317(n-i) != false {
			res = append(res, i)
			res = append(res, n-1)
			break
		}
	}
	return res
}
