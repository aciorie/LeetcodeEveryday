package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	// 创建拷贝节点并插入原链表
	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = newNode
		cur = newNode.Next
	}

	// 设置拷贝节点的 random 指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 分离原链表和拷贝链表
	dummyHead := &Node{}
	copyCurr := dummyHead
	origCurr := head

	for origCurr != nil {
		copyCurr.Next = origCurr.Next
		copyCurr = copyCurr.Next
		origCurr.Next = origCurr.Next.Next
		origCurr = origCurr.Next
	}

	return dummyHead.Next
}
