package codetop3rd

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1. 使用快慢指针找到链表的中点
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 断开链表，分成两半
	if prev != nil {
		prev.Next = nil
	}

	// 2. 递归地对左右两半进行排序
	left := sortList(head)
	right := sortList(slow)

	// 3. 合并两个已排序的子链表
	return merge_148(left, right)
}

// merge 函数：合并两个已排序的链表
func merge_148(l1, l2 *ListNode) *ListNode {
	// 创建一个虚拟头节点，简化合并逻辑
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	// 将剩余的节点连接到合并后的链表
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
