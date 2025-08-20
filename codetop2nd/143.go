package codetop2nd

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// 第一步：用快慢指针找到中点，并将链表一分为二
	mid := findMiddle(head)

	// 分割为两部分
	head1 := head
	head2 := mid.Next
	mid.Next = nil // 断开前后两部分链表

	// 第二步：反转后半部分
	head2 = reverseList_143(head2)

	// 第三步：交替合并两部分链表
	mergeLists(head1, head2)
}

// 辅助函数：使用快慢指针找到中点
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 辅助函数：反转链表
func reverseList_143(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

// 辅助函数：交替合并两个链表
func mergeLists(head1, head2 *ListNode) {
	var next1, next2 *ListNode
	for head1 != nil && head2 != nil {
		next1 = head1.Next
		next2 = head2.Next

		// head1 指向 head2
		head1.Next = head2
		// head2 指向原 head1 的下一个节点
		head2.Next = next1

		// 移动到下一个节点
		head1 = next1
		head2 = next2
	}
}
