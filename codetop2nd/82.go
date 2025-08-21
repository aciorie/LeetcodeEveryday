package codetop2nd

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		if prev.Next.Val == prev.Next.Next.Val {
			duplicateVal := prev.Next.Val
			for prev.Next != nil && prev.Next.Val == duplicateVal {
				prev.Next = prev.Next.Next
			}
		} else {
			prev = prev.Next
		}
	}
	return dummy.Next
}
