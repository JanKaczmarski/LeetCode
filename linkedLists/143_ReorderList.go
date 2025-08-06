package linkedlists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	// traverse to the middle of linked list
	// when in middle combine n // 2 - 1 with n // 2 and go
	s := make([]*ListNode, 0)
	n := 0
	start := head
	for head != nil {
		n++
		head = head.Next
	}
	head = start

	for i := 0; i < n/2; i++ {
		s = append(s, head)
		head = head.Next
	}

	var prev *ListNode

	// odd number of elements. The middle will point to nil and be the last one
	// in the resulting linked list
	// last element from new head has to point into tmp
	if n%2 == 1 {
		prev = head
		head = head.Next
		prev.Next = nil
	}

	for i := len(s) - 1; i >= 0; i-- {
		s[i].Next = head
		tmp := head.Next
		head.Next = prev
		head = tmp
		prev = s[i]
	}
}
