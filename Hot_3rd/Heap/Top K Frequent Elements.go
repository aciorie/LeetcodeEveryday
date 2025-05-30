package heap

import (
	"container/heap"
	"sort"
)

type Item struct {
	value     int
	frequency int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() any {
	tmp := *h
	top := tmp[len(tmp)-1]
	*h = tmp[:len(tmp)-1]
	return top
}

func topKFrequent(nums []int, k int) []int {
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
