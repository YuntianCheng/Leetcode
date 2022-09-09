#
# @lc app=leetcode.cn id=392 lang=python3
#
# [392] 判断子序列
#

# @lc code=start
class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if len(s) > len(t):
            return False

        if len(s) == len(t) and s != t:
            return False
        i = 0
        for t_char in t:
            if i < len(s):
                if s[i] == t_char:
                    i += 1
            else:
                break
        if i == len(s):
            return True
        return False
# @lc code=end

