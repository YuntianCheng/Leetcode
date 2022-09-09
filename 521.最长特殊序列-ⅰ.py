#
# @lc app=leetcode.cn id=521 lang=python3
#
# [521] 最长特殊序列 Ⅰ
#

# @lc code=start
class Solution:
    def findLUSlength(self, a: str, b: str) -> int:
        if a!=b:
            return len(a) if len(a) > len(b) else len(b)
        else:
            return -1
# @lc code=end

