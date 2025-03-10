package codinginterviews

/*
面试题 25. 合并两个排序的链表


题目描述
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode{
	dummy:=&ListNode{}
	cur:=dummy
	for l1!=nil&&l2!=nil{
		if l1.Val<l2.Val{
			cur.Next=l1
			l1=l1.Next
		}else{
			cur.Next=l2
			l2=l2.Next
		}
	}

	if l1==nil{
		cur.Next=l2
	}else{
		cur.Next=l1
	}

	return dummy.Next
}