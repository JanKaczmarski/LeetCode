package stack

type node struct {
	comb      string
	openCnt   int
	closedCnt int
}

func generateParenthesis(n int) []string {
	// take element from the stack. If len of the element is < n/2 then don't append.
	// else append. - HOw to make sure we generate correct combinations?
	stack := make([]node, 0)
	stack = append(stack, node{
		comb:      "(",
		openCnt:   1,
		closedCnt: 0,
	})
	solution := make([]string, 0)

	for len(stack) > 0 {
		lastNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Combination is invalid when this criteria is met
		if (lastNode.openCnt < lastNode.closedCnt) || (lastNode.openCnt > n) {
			continue
		}

		if len(lastNode.comb) == 2*n {
			solution = append(solution, lastNode.comb)
			continue
		}

		newClosed := node{
			comb:      lastNode.comb + ")",
			openCnt:   lastNode.openCnt,
			closedCnt: lastNode.closedCnt + 1,
		}
		newOpened := node{
			comb:      lastNode.comb + "(",
			openCnt:   lastNode.openCnt + 1,
			closedCnt: lastNode.closedCnt,
		}

		stack = append(stack, newClosed, newOpened)
	}

	return solution
}
