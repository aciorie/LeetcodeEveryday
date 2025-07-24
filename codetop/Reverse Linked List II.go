package codetop

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}

	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	start := pre.Next
	cur := start.Next

	for i := 0; i < right-left; i++ {
		tmp := cur.Next

		// 将cur插到pre后面
		cur.Next = pre.Next
		pre.Next = cur
		// 指向剩余未反转部分的头部
		start.Next = tmp
		// 移动到下一个等待反转的节点
		cur = tmp
	}
	return dummy.Next
}
