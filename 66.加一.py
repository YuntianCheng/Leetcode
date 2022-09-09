#
# @lc app=leetcode.cn id=66 lang=python3
#
# [66] 加一
#
from typing import List
# @lc code=start
class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        up = 1
        for i in range(len(digits) -1, -1, -1):
            if up == 1:
                if digits[i] < 9:
                    digits[i] += 1
                    up = 0
                else:
                    digits[i] = 0
                    up = 1
            if up == 0:
                return digits
        if up == 1:
            return [1] + digits
# @lc code=end
print(Solution().plusOne([8,9,9,9]))

