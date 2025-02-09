package codinginterviews

/*
面试题 06. 从尾到头打印链表


题目描述
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]


限制：

0 <= 链表长度 <= 10000
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func solution6_1(head *ListNode) (res []int) {
	for head.Next != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	for l, r := 0, len(res)-1; l < r; {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

// Recursive
// First recursively get the reverse value list of the nodes after the head, and then add the value of the head to the end of the list
func solution6_2(head *ListNode) (res []int) {
	if head == nil {
		return
	}
	res = solution6_2(head.Next)
	res = append(res, head.Val)
	return res
}
