package hot

/*Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Example 1:
Input: head = [1,1,2]
Output: [1,2]

Example 2:
Input: head = [1,1,2,3,3]
Output: [1,2,3]


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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	slow := dummy

	for fast := head; fast != nil; fast = fast.Next {
		//internal loop to get a different value from slow.Val
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}
		slow.Next = fast
		slow = slow.Next
	}
	return dummy.Next
}
