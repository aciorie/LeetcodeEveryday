package contest

/*
Q1. Maximum Subarray With Equal Products
Easy
3 pt.
You are given an array of positive integers nums.

An array arr is called product equivalent if prod(arr) == lcm(arr) * gcd(arr), where:

prod(arr) is the product of all elements of arr.
gcd(arr) is the GCD of all elements of arr.
lcm(arr) is the LCM of all elements of arr.
Return the length of the longest product equivalent subarray of nums.

A subarray is a contiguous non-empty sequence of elements within an array.

The term gcd(a, b) denotes the greatest common divisor of a and b.

The term lcm(a, b) denotes the least common multiple of a and b.



Example 1:
Input: nums = [1,2,1,2,1,1,1]
Output: 5
Explanation:
The longest product equivalent subarray is [1, 2, 1, 1, 1], where prod([1, 2, 1, 1, 1]) = 2, gcd([1, 2, 1, 1, 1]) = 1, and lcm([1, 2, 1, 1, 1]) = 2.

Example 2:
Input: nums = [2,3,4,5,6]
Output: 3
Explanation:
The longest product equivalent subarray is [3, 4, 5].

Example 3:
Input: nums = [1,2,3,1,4,5,1]
Output: 5



Constraints:

2 <= nums.length <= 100
1 <= nums[i] <= 10
*/

func MaxLength(nums []int) int {
	n := len(nums)
	res := 0

	for i := 0; i < n; i++ {
		prod := 1
		g := nums[i]
		l := nums[i]
		for j := i; j < n; j++ {
			prod *= nums[j]
			g = gcd(g, nums[j])
			l = lcm(l, nums[j])

			if prod == l*g {
				res = max(res, j-i+1)
			}
		}
	}

	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}
