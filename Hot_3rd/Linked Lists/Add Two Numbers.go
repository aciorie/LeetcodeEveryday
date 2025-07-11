package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry, cur := 0, dummy
	for l1 != nil || l2 != nil {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		digit := sum % 10
		cur.Next = &ListNode{Val: digit}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
