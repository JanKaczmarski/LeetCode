from typing import List


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        # go from left to right and then from right to left
        answer = []
        total = 1
        for num in nums:
            answer.append(total)
            total *= num

        total = 1
        for i in range(len(nums) - 1, -1, -1):
            answer[i] *= total
            total *= nums[i]

        return answer
