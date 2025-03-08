package mar

import "math"

/*
2379. Minimum Recolors to Get K Consecutive Black Blocks
Easy
Topics
Companies
Hint
You are given a 0-indexed string blocks of length n, where blocks[i] is either 'W' or 'B', representing the color of the ith block. The characters 'W' and 'B' denote the colors white and black, respectively.

You are also given an integer k, which is the desired number of consecutive black blocks.

In one operation, you can recolor a white block such that it becomes a black block.

Return the minimum number of operations needed such that there is at least one occurrence of k consecutive black blocks.



Example 1:

Input: blocks = "WBBWWBBWBW", k = 7
Output: 3
Explanation:
One way to achieve 7 consecutive black blocks is to recolor the 0th, 3rd, and 4th blocks
so that blocks = "BBBBBBBWBW".
It can be shown that there is no way to achieve 7 consecutive black blocks in less than 3 operations.
Therefore, we return 3.
Example 2:

Input: blocks = "WBWBBBW", k = 2
Output: 0
Explanation:
No changes need to be made, since 2 consecutive black blocks already exist.
Therefore, we return 0.


Constraints:

n == blocks.length
1 <= n <= 100
blocks[i] is either 'W' or 'B'.
1 <= k <= n
*/
func MinimumRecolors(blocks string, k int) int {
	n := len(blocks)
	dp := make([]int, n)

	for i := range dp {
		dp[i] = math.MaxInt32
	}

	// Count the number of white blocks at each position
	whiteNum := make([]int, n+1)
	for i := 0; i < n; i++ {
		whiteNum[i+1] = whiteNum[i] // Inherit the number of white blocks in the previous position
		if blocks[i] == 'W' {
			whiteNum[i+1]++
		}
	}

	// dp[i]=min(dp[i],dp[i−k]+number of white blocks that need to be turned black)
	for i := 0; i < n; i++ {
		if i >= k-1 {
			whiteInWindow := whiteNum[i+1] - whiteNum[i-k+1] // The number of white blocks in the current window
			dp[i] = min(dp[i], whiteInWindow)                // Update dp[i] to the number of white blocks in the current window
			if i > 0 {
				dp[i] = min(dp[i], dp[i-1])
			}
		}
	}

	// // If dp[n-1] is still the initial maximum value, it means that k consecutive black blocks cannot be formed
	if dp[n-1] == math.MaxInt32 {
		return 0
	}
	return dp[n-1]
}
