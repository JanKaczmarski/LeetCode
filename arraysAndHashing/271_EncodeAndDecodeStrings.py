from typing import List

DELIMIT = "#"


class Solution:

    def encode(self, strs: List[str]) -> str:
        """
        Encodes list of strings by using <count># delimter. <count> specifies how long
        next string is. Allows for decoding later on.

        Example:
        ["we","say"] -> "2#we3#say"
        """
        res = ""
        for s in strs:
            res += str(len(s)) + DELIMIT + s

        return res

    def decode(self, s: str) -> List[str]:
        """
        Decodes string in format <len(word)>#<word> to list element.

        Example:
        "2#we3#say" -> ["we", "say"]
        """
        if not s:
            return []
        res = []
        i = 0

        while i < len(s):
            jump_str = ""
            while i < len(s) and s[i] != DELIMIT:
                jump_str += s[i]
                i += 1

            jump = int(jump_str)
            res.append(s[i + 1 : i + jump + 1])
            i += jump + 1

        return res
