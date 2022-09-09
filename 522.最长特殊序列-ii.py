#
# @lc app=leetcode.cn id=522 lang=python3
#
# [522] 最长特殊序列 II
#
from typing import List

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
    def findLUSlength(self, strs: List[str]) -> int:
        max_length = -1
        length = len(strs)
        for i in range(length):
            isSubSeq = False
            cu = strs[i]
            for j in range(length):
                if i != j:
                    if self.isSubsequence(strs[i], strs[j]):
                        isSubSeq = True
                        break;
            if not isSubSeq:
                max_length = max(max_length, len(cu))
        return max_length

# @lc code=end

