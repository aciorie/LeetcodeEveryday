package codetop2nd

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	pre := head
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// curr 指向要反转的子列表的第一个节点
    curr := pre.Next
    
    // next 指针用于保存curr的下一个节点
    // 该循环将执行 right - left 次
    for i := 0; i < right - left; i++ {
        next := curr.Next
        
        // 1. curr.Next 跳过 next，指向 next 的下一个节点
        curr.Next = next.Next
        
        // 2. 将 next 指向 pre.Next (即curr)
        next.Next = pre.Next
        
        // 3. 将 pre 指向 next，完成头插
        pre.Next = next
    }

	return dummy.Next
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
    // 创建一个虚拟头节点
    dummy := &ListNode{Next: head}
    
    // 1. 找到要反转的子列表的前一个节点
    before_left := dummy
    for i := 0; i < left-1; i++ {
        before_left = before_left.Next
    }
    
    // sublist_head 是要反转的子列表的头节点 (2)
    sublist_head := before_left.Next
    
    // 2. 对子列表进行标准的链表反转
    var prev_rev *ListNode
    curr_rev := sublist_head
    
    for i := 0; i <= right-left; i++ {
        next_rev := curr_rev.Next
        curr_rev.Next = prev_rev
        prev_rev = curr_rev
        curr_rev = next_rev
    }
    
    // 此时，prev_rev 指向反转后子列表的新头部 (4)
    // sublist_head 仍指向原来的头部 (2)，但它的 Next 已经被修改
    // curr_rev 指向子列表之后的第一个节点 (5)
    
    // 3. 重新连接两端
    
    // 原来的 sublist_head (2) 现在是反转后的子列表的尾部
    // 它的 Next 指针应该指向子列表之后的节点 (5)
    sublist_head.Next = curr_rev
    
    // 将 before_left (1) 的 Next 指针指向反转后子列表的新头部 (4)
    before_left.Next = prev_rev
    
    return dummy.Next
}