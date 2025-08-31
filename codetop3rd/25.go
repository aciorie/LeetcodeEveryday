package codetop3rd

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	dummy := &ListNode{Next: head}
	prevGroupHead := dummy

	for {
		// check if there are at least k nodes
		count := 0
		cur := prevGroupHead.Next
		for count < k && cur != nil {
			count++
			cur = cur.Next
		}
		if count < k {
			break
		}

		// reverse k nodes
		prev := prevGroupHead.Next
		curr := prev.Next
		for i := 1; i < k; i++ {
			tmp := curr.Next
			tmp.Next = curr
			curr.Next = prev
			prevGroupHead.Next = curr
			cur = tmp
		}

		// Move prevGroupEnd to the end of the reversed group
		prevGroupHead = prev
	}
	return dummy.Next
}
