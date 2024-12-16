package Dec

import "container/heap"

/*
3264. Final Array State After K Multiplication Operations I		Easy

You are given an integer array nums, an integer k, and an integer multiplier.

You need to perform k operations on nums. In each operation:

Find the minimum value x in nums. If there are multiple occurrences of the minimum value, select the one that appears first.
Replace the selected minimum value x with x * multiplier.
Return an integer array denoting the final state of nums after performing all k operations.



Example 1:
Input: nums = [2,1,3,5,6], k = 5, multiplier = 2
Output: [8,4,6,5,6]
Explanation:
Operation	Result
After operation 1	[2, 2, 3, 5, 6]
After operation 2	[4, 2, 3, 5, 6]
After operation 3	[4, 4, 3, 5, 6]
After operation 4	[4, 4, 6, 5, 6]
After operation 5	[8, 4, 6, 5, 6]

Example 2:
Input: nums = [1,2], k = 3, multiplier = 4
Output: [16,8]
Explanation:
Operation	Result
After operation 1	[4, 2]
After operation 2	[4, 8]
After operation 3	[16, 8]


Constraints:
1 <= nums.length <= 100
1 <= nums[i] <= 100
1 <= k <= 10
1 <= multiplier <= 5
*/

type Item struct {
	value int
	index int
}

type minHeap []*Item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

//Wrong Answer		581 / 737 testcases passed		submitted at Dec 16, 2024 10:23
func GetFinalState1(nums []int, k int, multiplier int) []int {
	h := &minHeap{}
	heap.Init(h)

	for i, num := range nums {
		heap.Push(h, &Item{value: num, index: i})
	}

	for i := 0; i < k; i++ {
		minItem := heap.Pop(h).(*Item)
		minItem.value *= multiplier
		heap.Push(h, minItem)
	}

	res := make([]int, len(nums))
	for h.Len() > 0 {
		item := heap.Pop(h).(*Item)
		res[item.index] = item.value
	}
	return res
}

//followed on the solution page
type minHeap2 [][2]int

func (h minHeap2) Len() int          { return len(h) }
func (h minHeap2) Swap(i int, j int) { h[i], h[j] = h[j], h[i] }
func (h minHeap2) Less(i int, j int) bool {
	if h[i][1] == h[j][1] {
		return h[i][0] < h[j][0]
	}
	return h[i][1] < h[j][1]
}
func (h *minHeap2) Push(a interface{}) { *h = append(*h, a.([2]int)) }
func (h *minHeap2) Pop() interface{}   { l := len(*h); res := (*h)[l-1]; *h = (*h)[:l-1]; return res }

func GetFinalState(nums []int, k int, multiplier int) []int {
	mh := minHeap2{}
	heap.Init(&mh)
	for i := 0; i < len(nums); i++ {
		heap.Push(&mh, [2]int{i, nums[i]})
	}

	for k > 0 {
		mh[0][1] *= multiplier
		heap.Fix(&mh, 0)
		k--
	}

	for mh.Len() > 0 {
		v := heap.Pop(&mh).([2]int)
		nums[v[0]] = v[1]
	}
	return nums
}
