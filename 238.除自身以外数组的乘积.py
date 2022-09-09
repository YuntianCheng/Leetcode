#
# @lc app=leetcode.cn id=238 lang=python3
#
# [238] 除自身以外数组的乘积
#
from os import stat
from typing import List
# @lc code=start
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        result = [1]*len(nums)
        start = 1
        end = 1
        for i in range(len(nums) - 1):
            start *= nums[i]
            end *= nums[len(nums) - 1 - i]
            result[i + 1] *= start
            result[len(nums) - 2 - i] *= end
        return result
# @lc code=end

