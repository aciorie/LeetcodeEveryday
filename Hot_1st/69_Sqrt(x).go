package hot

// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

// You must not use any built-in exponent function or operator.

// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

// Example 1:

// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.
// Example 2:

// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

// Constraints:

// 0 <= x <= 2^31 - 1

//dumb method,iterate from 2 to x
func MySqrt(x int) int {
	if x < 2 {
		return x
	} else if x < 4 {
		return 1
	}

	res := 0
	for i := 2; i < x; i++ {
		if i*i == x {
			res = i
			break
		} else if i*i > x {
			res = i - 1
			break
		}
		continue
	}
	return res
}

//a better method,using binary search
func MySqrt2(x int) int {
	if x < 2 {
		return x
	}
	l, r := 0, x
	for l < r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}

/* this method definitely better than my original thoughts.
i was considering some strategies like to calculate the odd number,
which using math.Sqrt() would be a good way
*/
