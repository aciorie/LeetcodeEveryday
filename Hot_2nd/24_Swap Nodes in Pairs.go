package hot2nd

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	// At least two nodes can be exchanged
	for prev.Next != nil && prev.Next.Next != nil {
		fir, sec := prev.Next, prev.Next.Next

		fir.Next = sec.Next
		sec.Next = fir
		prev.Next = sec
		// Move the prev pointer to the end of the linked list that has been swapped
		prev = fir
	}
	return dummy.Next
}
