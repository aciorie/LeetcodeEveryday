package review

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	oldToNew := make(map[*Node]*Node)

	curr := head
	for curr != nil {
		oldToNew[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		copiedNode := oldToNew[curr]
		copiedNode.Next = oldToNew[curr.Next]
		copiedNode.Random = oldToNew[curr.Random]
		curr = curr.Next
	}

	return oldToNew[head]
}
