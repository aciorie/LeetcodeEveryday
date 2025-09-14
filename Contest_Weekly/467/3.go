package Contest

import "sort"

func subsequenceSumAfterCapping(nums []int, k int) []bool {
	n := len(nums)
	answer := make([]bool, n)

	// Sort the nums array to process capped values efficiently
	sort.Ints(nums)

	// dp[i] is true if sum i is possible
	dp := make([]bool, k+1)
	dp[0] = true // Sum of 0 is always possible

	numIdx := 0

	// We iterate through x from 1 to n.
	// However, the set of available numbers only changes when x becomes equal to a number in the sorted nums array.
	// So we can process x values in blocks.
	currentX := 1
	for numIdx < n {
		// Process x values from the previous number up to the current number
		for currentX < nums[numIdx] && currentX <= n {
			answer[currentX-1] = dp[k]
			currentX++
		}

		if currentX > n {
			break
		}

		// Update DP table for all elements equal to the current number
		for numIdx < n && nums[numIdx] == currentX {
			for sum := k; sum >= currentX; sum-- {
				dp[sum] = dp[sum] || dp[sum-currentX]
			}
			numIdx++
		}

		// The answer for x = currentX is the state of the DP table after updates
		answer[currentX-1] = dp[k]
		currentX++
	}

	// Handle any remaining x values that are greater than the largest number in nums
	for currentX <= n {
		answer[currentX-1] = dp[k]
		currentX++
	}

	return answer
}
