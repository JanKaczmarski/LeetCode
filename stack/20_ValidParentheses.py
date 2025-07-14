class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        for par in s:
            if par in ("(", "[", "{"):
                stack.append(par)

            # closing bracket and empty stack
            elif not stack:
                return False
            else:
                if (
                    (par == ")" and stack[-1] != "(")
                    or (par == "]" and stack[-1] != "[")
                    or (par == "}" and stack[-1] != "{")
                ):
                    return False
                else:
                    stack.pop()

        return len(stack) == 0
