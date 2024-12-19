package Dec

/*
769. Max Chunks To Make Sorted		Attempted		Medium

You are given an integer array arr of length n that represents a permutation of the integers in the range [0, n - 1].

We split arr into some number of chunks (i.e., partitions), and individually sort each chunk. After concatenating them, the result should equal the sorted array.

Return the largest number of chunks we can make to sort the array.



Example 1:
Input: arr = [4,3,2,1,0]
Output: 1
Explanation:
Splitting into two or more chunks will not return the required result.
For example, splitting into [4, 3], [2, 1, 0] will result in [3, 4, 0, 1, 2], which isn't sorted.

Example 2:
Input: arr = [1,0,2,3,4]
Output: 4
Explanation:
We can split into two chunks, such as [1, 0], [2, 3, 4].
However, splitting into [1, 0], [2], [3], [4] is the highest number of chunks possible.


Constraints:

n == arr.length
1 <= n <= 10
0 <= arr[i] < n
All the elements of arr are unique.
*/

//2 failed attempts
func maxChunksToSorted1(arr []int) int {
	if arr[len(arr)-1] == 0 {
		return 1
	}
	res, n := 0, len(arr)

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			res = n - i
		}
	}

	return res
}

func maxChunksToSorted2(arr []int) int {
	if arr[len(arr)-1] == 0 {
		return 1
	}
	res, n := 0, len(arr)-1
	stack := make([]int, n)
	stack[0] = 0

	for i := 1; i < n; i++ {
		stack[i] = formula(i)
		arr[i] += arr[i-1]
		if arr[i]-stack[i] <= 0 {
			res++
		}
	}

	return res
}

func formula(a int) int {
	return a * (a + 1) / 2
}

func MaxChunksToSorted(arr []int) int {
	maxValue := 0
	chunks := 0

	for i, value := range arr {
		// Update the maximum value seen so far
		if value > maxValue {
			maxValue = value
		}
		// If the maximum value equals the current index, we can form a chunk
		if maxValue == i {
			chunks++
		}
	}

	return chunks
}
