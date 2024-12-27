package Dec

/*
1014. Best Sightseeing Pair
Medium
Topics
Companies
Hint
You are given an integer array values where values[i] represents the value of the ith sightseeing spot. Two sightseeing spots i and j have a distance j - i between them.

The score of a pair (i < j) of sightseeing spots is values[i] + values[j] + i - j: the sum of the values of the sightseeing spots, minus the distance between them.

Return the maximum score of a pair of sightseeing spots.



Example 1:
Input: values = [8,1,5,2,6]
Output: 11
Explanation: i = 0, j = 2, values[i] + values[j] + i - j = 8 + 5 + 0 - 2 = 11

Example 2:
Input: values = [1,2]
Output: 2


Constraints:

2 <= values.length <= 5 * 104
1 <= values[i] <= 1000
*/

//brute force,got Time Limit Exceeded	50 / 55 testcases passed
func maxScoreSightseeingPair1(values []int) int {
	if len(values) == 2 {
		return values[0] + values[1] - 1
	}

	res := 0

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			res = max(res, values[i]+values[j]+i-j)
		}
	}
	return res
}

//try another way
func MaxScoreSightseeingPair(values []int) int {
	if len(values) == 2 {
		return values[0] + values[1] - 1
	}

	res := 0
	//Store the middle value of the current maximum score
	maxI := values[0]

	for j := 1; j < len(values); j++ {
		score := maxI + values[j] - j
		res = max(res, score)

		maxI = max(maxI, values[j]+j)
	}

	return res
}
