#
# @lc app=leetcode.cn id=434 lang=python3
#
# [434] 字符串中的单词数
#

# @lc code=start
class Solution:
    def countSegments(self, s: str) -> int:
        count = 0
        strs = s.split(' ')
        for word in strs:
            r = word.strip()
            if len(r) > 0:
                count += 1
        return count
# @lc code=end

