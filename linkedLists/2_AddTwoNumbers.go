package linkedlists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// add two nodes, each one of different input list
	// use `carry` to store value that exceded 10 -> for 5 + 7 carry will be 2

	carry := 0
	head := &ListNode{}
	res := head

	// value stored in current iteration
	var val int
	for l1 != nil && l2 != nil {
		// compute result node value and carry
		val = l1.Val + l2.Val + carry
		carry = 0
		if val >= 10 {
			carry = 1
			val = val - 10
		}

		// New result node creation
		newNode := &ListNode{
			Val: val,
		}
		head.Next = newNode
		head = head.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		val = l1.Val + carry
		carry = 0
		if val >= 10 {
			carry = 1
			val = val - 10
		}

		newNode := &ListNode{
			Val: val,
		}
		head.Next = newNode
		head = head.Next

		l1 = l1.Next
	}

	for l2 != nil {
		val = l2.Val + carry
		carry = 0
		if val >= 10 {
			carry = 1
			val = val - 10
		}

		newNode := &ListNode{
			Val: val,
		}
		head.Next = newNode
		head = head.Next

		l2 = l2.Next
	}

	if carry == 1 {
		newNode := &ListNode{
			Val: carry,
		}
		head.Next = newNode
		head = head.Next
	}

	return res.Next
}
