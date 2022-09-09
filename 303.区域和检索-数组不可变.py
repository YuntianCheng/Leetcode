#
# @lc app=leetcode.cn id=303 lang=python3
#
# [303] 区域和检索 - 数组不可变
#
from typing import List
# @lc code=start
class NumArray:

    def __init__(self, nums: List[int]):
        self.nums = nums
        sum = 0
        for i in range(len(nums)):
            sum = sum + nums[i]
            self.nums[i] = sum


    def sumRange(self, left: int, right: int) -> int:
        if left == 0:
            return self.nums[right]
        return self.nums[right] - self.nums[left-1]


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)
# @lc code=end

n = NumArray([-2,0,3,-5,2,-1])
print(n.sumRange(2,5))