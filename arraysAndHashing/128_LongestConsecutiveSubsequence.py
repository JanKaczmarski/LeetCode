from typing import List, Tuple


# This one was my solution and idea. It's not the most optimal solution, because using
# set in the solution below is much faster, but this solution is still pretty neet
class SolutionFindUnion:
    def longestConsecutive(self, nums: List[int]) -> int:
        """
        Solutions that uses find-union to represent a chain with it's representative. The map is used
        to store chain elements.
        """

        # base case for find-union structure
        # num: <chain_len, chain representant)
        store = {num: (1, num) for num in nums}

        def find(x: int) -> Tuple[int, int]:
            """
            For given x it finds representant of the chain that x is part of. It does so recursively, updating
            all memeber of the chain that x (and next elements) were pointing to
            Retruns Tuple[<chain_length>, <chain_representant>]
            """
            if x == store[x][1]:
                return store[x]

            store[x] = find(store[x][1])
            return store[x]

        def union(x: int, y: int) -> None:
            """
            Combines x and y into single chain with common representant. New chain length is length of chain
            that x is a part of and len of chain that y is a part of. New chain representant will always be x chain
            representant. Operates on `store` dictionary that is initialized in outer scope.
            """
            x_repr = find(x)
            y_repr = find(y)
            if x_repr[1] == y_repr[1]:
                return

            new_length = x_repr[0] + y_repr[0]
            store[x_repr[1]] = (new_length, x_repr[1])
            store[y_repr[1]] = (new_length, x_repr[1])

        for num in nums:
            if num - 1 in store:
                union(num, num - 1)
            if num + 1 in store:
                union(num, num + 1)

        sol = 0
        for _, (length, _) in store.items():
            sol = max(sol, length)

        return sol


# straightforward solution, It works because of if statement that ensure
# we don't recomupte a chain length for each of it's elements
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        sol = 0
        nums_s = set(nums)

        for num in nums_s:
            if num - 1 not in nums_s:
                length = 1
                while (num + length) in nums_s:
                    length += 1
                sol = max(sol, length)

        return sol
