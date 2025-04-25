package apr

/*
1399. Count Largest Group
Easy
Topics
Companies
Hint
You are given an integer n.

Each number from 1 to n is grouped according to the sum of its digits.

Return the number of groups that have the largest size.



Example 1:

Input: n = 13
Output: 4
Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9].
There are 4 groups with largest size.
Example 2:

Input: n = 2
Output: 2
Explanation: There are 2 groups [1], [2] of size 1.


Constraints:

1 <= n <= 104
*/
func CountLargestGroup(n int) int {
	if n == 1 {
		return 1
	} else if n < 10 {
		return n
	}
	sum := make(map[int]int)
	res, largestSize := 0, 0
	for i := 0; i <= n; i++ {
		digitsSum := digitSum(i)
		sum[digitsSum]++
		if largestSize < sum[digitsSum] {
			largestSize = sum[digitsSum]
		}
	}
	for _, v := range sum {
		if v == largestSize {
			res++
		}
	}
	return res
}

func digitSum(i int) int {
	if i < 10 {
		return i
	}
	res := 0
	for i > 0 {
		res += (i % 10)
		i /= 10
	}
	return res
}
