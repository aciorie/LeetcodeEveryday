package codetop3rd

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	valInMap := make(map[int]int)
	for i := head; i != nil; i = i.Next {
		valInMap[i.Val]++
	}

	newHead := &ListNode{}
	res := newHead
	for cur := head; cur != nil; cur = cur.Next {
		if valInMap[cur.Val] == 1 {
			res.Next = &ListNode{Val: cur.Val}
			res = res.Next
		}
	}
	return newHead.Next
}
