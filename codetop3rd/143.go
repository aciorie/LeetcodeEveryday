package codetop3rd

func reorderList(head *ListNode) {
	// find the middle position
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// break the first half
	secondHalf := slow.Next
	slow.Next = nil

	// reverse the second half
	var prev *ListNode
	curr := secondHalf
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	secondHalf = prev

	// iterate both lists
	firstHalf := head
	for firstHalf != nil && secondHalf != nil {
		// 保存后半部分和前半部分的下一个节点
		temp1 := firstHalf.Next
		temp2 := secondHalf.Next

		// 将前半部分的当前节点连接到后半部分的当前节点
		firstHalf.Next = secondHalf
		// 将后半部分的当前节点连接到前半部分的下一个节点
		secondHalf.Next = temp1

		// 移动指针到下一个位置
		firstHalf = temp1
		secondHalf = temp2
	}
}
