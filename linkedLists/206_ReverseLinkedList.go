package linkedlists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	if head != nil {
		head.Next = prev
	}
	return head
}
