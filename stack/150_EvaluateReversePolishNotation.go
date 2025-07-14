package stack

import (
	"slices"
	"strconv"
)

func evalRPN(tokens []string) int {
	operators := []string{"+", "-", "*", "/"}
	stack := make([]int, 0)

	for _, token := range tokens {
		if slices.Contains(operators, token) {
			var res int
			firstOperand := stack[len(stack)-2]
			secondOperand := stack[len(stack)-1]

			switch token {
			case "+":
				res = firstOperand + secondOperand
			case "-":
				res = firstOperand - secondOperand
			case "*":
				res = firstOperand * secondOperand
			case "/":
				res = firstOperand / secondOperand
			}

			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		} else {
			tokenInt, _ := strconv.Atoi(token)

			stack = append(stack, tokenInt)
		}
	}

	return stack[0]
}
