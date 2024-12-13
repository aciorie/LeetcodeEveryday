package hot

/*
Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
leaving only distinct numbers from the original list. Return the linked list sorted as well.

Example 1:
Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

Example 2:
Input: head = [1,1,1,2,3]
Output: [2,3]


Constraints:
The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//can be passed on the leetcode website,but failed at local tests
//Runtime	2ms		Beats 7.84%		Memory	5.34MB	Beats 6.59%
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	//count how many times each value appears
	nodesInMap := make(map[int]int)
	for cur := head; cur != nil; cur = cur.Next {
		nodesInMap[cur.Val]++
	}

	//iterate the hash table, if one's value = 1, pass it to the new linked list
	newHead := &ListNode{}
	res := newHead
	for cur := head; cur != nil; cur = cur.Next {
		if nodesInMap[cur.Val] == 1 {
			res.Next = &ListNode{Val: cur.Val}
			res = res.Next
		}
	}
	return newHead.Next
}

//copied from one of the solutions
func deleteDuplicates2(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}
		pre, cur = cur, cur.Next
	}

	return dummy.Next
}
