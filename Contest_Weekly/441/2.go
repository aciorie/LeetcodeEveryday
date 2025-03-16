package contest

/*
Q2. Closest Equal Element Queries
Medium
4 pt.
You are given a circular array nums and an array queries.

For each query i, you have to find the following:

The minimum distance between the element at index queries[i] and any other index j in the circular array, where nums[j] == nums[queries[i]]. If no such index exists, the answer for that query should be -1.
Return an array answer of the same size as queries, where answer[i] represents the result for query i.



Example 1:

Input: nums = [1,3,1,4,1,3,2], queries = [0,3,5]

Output: [2,-1,3]

Explanation:

Query 0: The element at queries[0] = 0 is nums[0] = 1. The nearest index with the same value is 2, and the distance between them is 2.
Query 1: The element at queries[1] = 3 is nums[3] = 4. No other index contains 4, so the result is -1.
Query 2: The element at queries[2] = 5 is nums[5] = 3. The nearest index with the same value is 1, and the distance between them is 3 (following the circular path: 5 -> 6 -> 0 -> 1).
Example 2:

Input: nums = [1,2,3,4], queries = [0,1,2,3]

Output: [-1,-1,-1,-1]

Explanation:

Each value in nums is unique, so no index shares the same value as the queried element. This results in -1 for all queries.



Constraints:

1 <= queries.length <= nums.length <= 105
1 <= nums[i] <= 106
0 <= queries[i] < nums.length
*/

// Time Limit Exceeded		612 / 614 testcases passed
func SolveQueries(nums []int, queries []int) []int {
	n := len(nums)
	indices := make(map[int][]int) // Value -> List of indices
	for i, num := range nums {
		indices[num] = append(indices[num], i)
	}

	answer := make([]int, len(queries))
	for i, queryIndex := range queries {
		targetValue := nums[queryIndex]
		if len(indices[targetValue]) <= 1 {
			answer[i] = -1 // No other occurrence
			continue
		}

		minDist := n // Initialize with a large distance
		for _, otherIndex := range indices[targetValue] {
			if otherIndex == queryIndex {
				continue // Skip the same index
			}
			dist1 := abs(otherIndex - queryIndex)
			dist2 := n - dist1
			minDist = min(minDist, min(dist1, dist2))
		}
		answer[i] = minDist
	}

	return answer
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
