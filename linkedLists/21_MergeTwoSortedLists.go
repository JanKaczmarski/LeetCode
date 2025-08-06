package linkedlists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	guard := &ListNode{}
	start := guard

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			guard.Next = list1
			list1 = list1.Next
		} else {
			guard.Next = list2
			list2 = list2.Next
		}
		guard = guard.Next
	}

	for list1 != nil {
		guard.Next = list1
		list1 = list1.Next
		guard = guard.Next
	}
	for list2 != nil {
		guard.Next = list2
		list2 = list2.Next
		guard = guard.Next
	}

	return start.Next
}
