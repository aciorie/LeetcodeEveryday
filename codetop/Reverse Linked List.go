package codetop

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	prev, cur := dummy, head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return cur
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return dummy
}
