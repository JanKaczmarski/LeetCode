package stack

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	minVal := math.MaxInt32
	if len(this.minStack) > 0 {
		minVal = this.minStack[len(this.minStack)-1]
	}
	if val < minVal {
		minVal = val
	}

	this.minStack = append(this.minStack, minVal)
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.minStack = this.minStack[:len(this.minStack)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
