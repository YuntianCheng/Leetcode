#
# @lc app=leetcode.cn id=553 lang=python3
#
# [553] 最优除法
#
from typing import List
# @lc code=start
class Solution:
    def optimalDivision(self, nums: List[int]) -> str:
        if len(nums) == 1:
            return str(nums[0])
        if len(nums) == 2:
            return '/'.join([str(num) for num in nums])
        nums_str = [str(num) for num in nums]
        nums_str[1] = '(' + nums_str[1]
        nums_str[-1] = nums_str[-1] + ')'
        return '/'.join(nums_str)
# @lc code=end
s = Solution()
print(s.optimalDivision([1,2,3,3]))
