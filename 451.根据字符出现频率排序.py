#
# @lc app=leetcode.cn id=451 lang=python3
#
# [451] 根据字符出现频率排序
#

# @lc code=start
class Solution:
    def frequencySort(self, s: str) -> str:
        nums = {}
        for c in s:
            if c in nums:
                nums[c] += 1
            else:
                nums[c] = 1
        s_list = list(nums.items())
        s_list.sort(key=lambda item: item[1], reverse=True)
        return ''.join(map(lambda item:item[0]*item[1], s_list))
# @lc code=end

