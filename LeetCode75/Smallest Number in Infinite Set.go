package leetcode75

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	top := old[len(old)-1]
	*h = old[:len(old)-1]
	return top
}

type SmallestInfiniteSet struct {
	nextNum int
	heap    IntHeap
	set     map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		nextNum: 1,
		heap:    IntHeap{},
		set:     make(map[int]bool),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.heap) > 0 {
		smallest := heap.Pop(&this.heap).(int)
		this.set[smallest] = false
		return smallest
	}
	smallest := this.nextNum
	this.nextNum++
	return smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.set[num] == true {
		return
	}
	if num >= this.nextNum {
		return
	}

	this.set[num] = true
	heap.Push(&this.heap, num)
}
