package linkedlists

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		fir, sec := prev.Next, prev.Next.Next

		fir.Next = sec.Next
		sec.Next = fir
		prev.Next = sec

		prev = fir
	}
	return dummy.Next
}
