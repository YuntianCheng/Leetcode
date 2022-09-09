#
# @lc app=leetcode.cn id=383 lang=python3
#
# [383] 赎金信
#

# @lc code=start


class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        m = {}
        for c in magazine:
            if c in m:
                m[c] += 1
            else:
                m[c] = 1
        for c in ransomNote:
            if c not in m or m[c] == 0:
                return False
            else:
                m[c] -= 1
        return True
# @lc code=end

