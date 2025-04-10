package codinginterviews

import (
	"container/heap"
	"sort"
)

/*
面试题 41. 数据流中的中位数


题目描述
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例 1：

输入：
["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]
输出：[null,null,null,1.50000,null,2.00000]
示例 2：

输入：
["MedianFinder","addNum","findMedian","addNum","findMedian"]
[[],[2],[],[3],[]]
输出：[null,null,2.00000,null,2.50000]


限制：

最多会对 addNum、findMedian 进行 50000 次调用。
*/

// Method 1
// not the best solution
type MedianFinder_1 struct {
	nums []int
}

func Constructor_41_1() MedianFinder_1 {
	return MedianFinder_1{nums: []int{}}
}

func (this *MedianFinder_1) AddNum_1(num int) {
	this.nums = append(this.nums, num)
}

func (this *MedianFinder_1) FindMedian_1() float64 {
	sort.Ints(this.nums)
	n := len(this.nums)
	if n%2 == 0 {
		return float64(this.nums[n/2-1]+this.nums[n/2]) / 2.0
	}
	return float64(this.nums[n/2])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// Method 2
// Time Limit Exceeded		20 / 21 testcases passed
type MedianFinder_2 struct {
	nums []int
}

func Constructor_41_2() MedianFinder_2 {
	return MedianFinder_2{nums: []int{}}
}

func (this *MedianFinder_2) AddNum_2(num int) {
	this.nums = append(this.nums, num)
	sort.Ints(this.nums)
}

func (this *MedianFinder_2) FindMedian_2() float64 {
	n := len(this.nums)
	if n%2 == 0 {
		return float64(this.nums[n/2-1]+this.nums[n/2]) / 2.0
	}
	return float64(this.nums[n/2])
}

// Method 3
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	maxHeap MaxHeap
	minHeap MinHeap
}

func Constructor_41() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 || num < this.maxHeap[0] {
		heap.Push(&this.maxHeap, num)
	} else {
		heap.Push(&this.minHeap, num)
	}

	// Balance heaps
	if this.maxHeap.Len()-this.minHeap.Len() > 1 {
		top := heap.Pop(&this.maxHeap).(int)
		heap.Push(&this.minHeap, top)
	} else if this.minHeap.Len()-this.maxHeap.Len() > 1 {
		top := heap.Pop(&this.minHeap).(int)
		heap.Push(&this.maxHeap, top)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		if this.maxHeap.Len() == 0 {
			return 0.0
		}
		return float64(this.maxHeap[0]+this.minHeap[0]) / 2.0
	} else if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(this.maxHeap[0])
	} else {
		return float64(this.minHeap[0])
	}
}
