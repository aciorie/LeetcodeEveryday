package Dec

import "sort"

/*
2593. Find Score of an Array After Marking All Elements		Medium

You are given an array nums consisting of positive integers.

Starting with score = 0, apply the following algorithm:

Choose the smallest integer of the array that is not marked. If there is a tie, choose the one with the smallest index.
Add the value of the chosen integer to score.
Mark the chosen element and its two adjacent elements if they exist.
Repeat until all the array elements are marked.
Return the score you get after applying the above algorithm.



Example 1:
Input: nums = [2,1,3,4,5,2]
Output: 7
Explanation: We mark the elements as follows:
- 1 is the smallest unmarked element, so we mark it and its two adjacent elements: [2,1,3,4,5,2].
- 2 is the smallest unmarked element, so we mark it and its left adjacent element: [2,1,3,4,5,2].
- 4 is the only remaining unmarked element, so we mark it: [2,1,3,4,5,2].
Our score is 1 + 2 + 4 = 7.

Example 2:
Input: nums = [2,3,5,1,3,2]
Output: 5
Explanation: We mark the elements as follows:
- 1 is the smallest unmarked element, so we mark it and its two adjacent elements: [2,3,5,1,3,2].
- 2 is the smallest unmarked element, since there are two of them, we choose the left-most one, so we mark the one at index 0 and its right adjacent element: [2,3,5,1,3,2].
- 2 is the only remaining unmarked element, so we mark it: [2,3,5,1,3,2].
Our score is 1 + 2 + 2 = 5.


Constraints:

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^6
*/

func FindScore(nums []int) int64 {
	n := len(nums)
	marked := make([]bool, n) // To track marked elements
	type element struct {
		value int
		index int
	}

	elements := make([]element, n)
	for i, num := range nums {
		elements[i] = element{value: num, index: i}
	}

	// Sort the elements by value and by index in case of ties
	sort.Slice(elements, func(i, j int) bool {
		if elements[i].value == elements[j].value {
			return elements[i].index < elements[j].index
		}
		return elements[i].value < elements[j].value
	})

	res := int64(0)
	for _, e := range elements {
		if !marked[e.index] { // If the element is not marked
			res += int64(e.value)
			// Mark this element and its adjacent elements if they exist
			marked[e.index] = true
			if e.index > 0 {
				marked[e.index-1] = true
			}
			if e.index < n-1 {
				marked[e.index+1] = true
			}
		}
	}

	return res
}
