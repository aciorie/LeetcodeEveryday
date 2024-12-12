package Dec

import (
	"container/heap"
	"math"
	"sort"
)

/*
2558. Take Gifts From the Richest Pile		Easy

You are given an integer array gifts denoting the number of gifts in various piles. Every second, you do the following:

Choose the pile with the maximum number of gifts.
If there is more than one pile with the maximum number of gifts, choose any.
Leave behind the floor of the square root of the number of gifts in the pile. Take the rest of the gifts.
Return the number of gifts remaining after k seconds.



Example 1:
Input: gifts = [25,64,9,4,100], k = 4
Output: 29
Explanation:
The gifts are taken in the following way:
- In the first second, the last pile is chosen and 10 gifts are left behind.
- Then the second pile is chosen and 8 gifts are left behind.
- After that the first pile is chosen and 5 gifts are left behind.
- Finally, the last pile is chosen again and 3 gifts are left behind.
The final remaining gifts are [5,8,9,4,3], so the total number of gifts remaining is 29.

Example 2:
Input: gifts = [1,1,1,1], k = 4
Output: 4
Explanation:
In this case, regardless which pile you choose, you have to leave behind 1 gift in each pile.
That is, you can't take any pile with you.
So, the total gifts remaining are 4.


Constraints:

1 <= gifts.length <= 103
1 <= gifts[i] <= 109
1 <= k <= 103
*/

// original one in Sep
func pickGifts(gifts []int, k int) int64 {
	sort.Ints(gifts)
	var res int64
	lastId := len(gifts) - 1
	for i := 0; i < k; i++ {
		sqrt := math.Sqrt(float64(gifts[lastId]))
		gifts[lastId] = int(sqrt)
		sort.Ints(gifts)
	}
	for _, val := range gifts {
		res += int64(val)
	}
	return res
}

// use heap

/*
Wrong Answer	52 / 102 testcases passed

Input	gifts =	[54,6,34,66,63,52,39,62,46,75,28,65,18,37,18,13,33,69,19,40,13,10,43,61,72]
		k = 7

Output	707
Expected	618
*/

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(i any)        { *h = append(*h, i.(int)) }
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func PickGifts(gifts []int, k int) int64 {
	h := &Heap{}
	heap.Init(h)

	for _, gift := range gifts {
		h.Push(gift)
	}

	for i := 0; i < k; i++ {
		gift := heap.Pop(h).(int)
		sqrt := math.Sqrt(float64(gift))
		remain := int(sqrt)
		h.Push(remain)
	}

	res := int64(0)
	for h.Len() > 0 {
		gift := h.Pop().(int)
		res += int64(gift)
	}
	return res
}
