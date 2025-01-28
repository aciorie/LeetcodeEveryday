package hot

import (
	"container/heap"
	"sort"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.



Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]


Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.


Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/

type Item struct {
	value     int
	frequency int
}

type MinHeap []Item

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].frequency < h[j].frequency }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func TopKFrequent(nums []int, k int) []int {
	ma := make(map[int]int)
	for _, num := range nums {
		ma[num]++
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for v, f := range ma {
		heap.Push(minHeap, Item{value: v, frequency: f})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	res := make([]int, 0, k)
	for minHeap.Len() > 0 {
		item := heap.Pop(minHeap).(Item)
		res = append(res, item.value)
	}
	sort.Slice(res, func(i, j int) bool {
		return ma[res[i]] > ma[res[j]]
	})
	return res
}
