package mar

import "math"

/*
2523. Closest Prime Numbers in Range
Medium
Topics
Companies
Hint
Given two positive integers left and right, find the two integers num1 and num2 such that:

left <= num1 < num2 <= right .
Both num1 and num2 are prime numbers.
num2 - num1 is the minimum amongst all other pairs satisfying the above conditions.
Return the positive integer array ans = [num1, num2]. If there are multiple pairs satisfying these conditions, return the one with the smallest num1 value. If no such numbers exist, return [-1, -1].



Example 1:

Input: left = 10, right = 19
Output: [11,13]
Explanation: The prime numbers between 10 and 19 are 11, 13, 17, and 19.
The closest gap between any pair is 2, which can be achieved by [11,13] or [17,19].
Since 11 is smaller than 17, we return the first pair.
Example 2:

Input: left = 4, right = 6
Output: [-1,-1]
Explanation: There exists only one prime number in the given range, so the conditions cannot be satisfied.


Constraints:

1 <= left <= right <= 106
*/

//Wrong Answer	42 / 66 testcases passed
func ClosestPrimes(left int, right int) []int {
	res := make([]int, 0)
	if left%2 == 0 {
		left++
	}
	if left == right {
		return []int{-1, -1}
	}
	for ; left <= right; left += 2 {
		if len(res) == 2 {
			break
		}
		if isPrime(left) {
			res = append(res, left)
		}
	}
	if len(res) != 2 {
		return []int{-1, -1}
	}
	return res
}

func isPrime(left int) bool {
	for i := 2; i <= int(math.Sqrt(float64(left))); i++ {
		if left%i == 0 {
			return false
		}
	}
	return true
}

func ClosestPrimes2(left int, right int) []int {
	minDiff, res, lastPrime := math.MaxInt32, []int{-1, -1}, -1

	isPrime := make([]bool, right+1)
	for i := 2; i <= right; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= right; i++ {
		if isPrime[i] {
			for j := i * i; j <= right; j += i {
				isPrime[j] = false
			}
		}
	}

	for num := left; num <= right; num++ {
		if isPrime[num] {
			if lastPrime != -1 {
				diff := num - lastPrime
				if diff < minDiff {
					minDiff = diff
					res[0] = lastPrime
					res[1] = num
				} else if diff == minDiff && (res[0] == -1 || lastPrime < res[0]) {
					res[0] = lastPrime
					res[1] = num
				}
			}
			lastPrime = num
		}
	}
	return res
}
