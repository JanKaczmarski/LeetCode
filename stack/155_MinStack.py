class MinStack:
    def __init__(self):
        # holds all values on stack like regular stack would
        self.stack = []
        # holds minimum value on a `stack` up to this point in stack
        # in other words each node in stack has a minimum element assigned to it
        self.min_stack = []

    def push(self, val: int) -> None:
        if self.min_stack:
            min_val = self.min_stack[-1]
        else:
            min_val = float("inf")

        min_val = min(min_val, val)

        self.stack.append(val)
        self.min_stack.append(min_val)

    def pop(self) -> None:
        self.stack.pop()
        self.min_stack.pop()

    def top(self) -> int:
        return self.stack[-1]

    def getMin(self) -> int:
        return self.min_stack[-1]


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
