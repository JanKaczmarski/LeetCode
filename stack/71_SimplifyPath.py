class Solution:
    def simplifyPath(self, path: str) -> str:
        """
        The idea is that when meet .. directory on our path, we simply pop from our stack of
        direcotries. So stack keeps track of all current direcotreis, and in some special case
        we make behavior different (./ ../ and /)
        """
        stack = []
        cur = ""

        for c in path + "/":
            if c == "/":
                if cur == "..":
                    if stack:
                        stack.pop()
                elif cur != "" and cur != ".":
                    stack.append(cur)
                cur = ""
            else:
                cur += c

        return "/" + "/".join(stack)
