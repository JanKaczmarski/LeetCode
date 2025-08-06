package linkedlists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// prev will be used to "remove" slow Node which is our n-th target
	// when fast reaches nil
	prev := &ListNode{
		Next: slow,
	}
	start := prev

	// until fast pointer isn't the last element of linked list
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}

	prev.Next = slow.Next

	return start.Next
}
