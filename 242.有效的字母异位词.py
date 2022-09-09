#
# @lc app=leetcode.cn id=242 lang=python3
#
# [242] 有效的字母异位词
#

# @lc code=start
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        s_statistics = {}
        for char in s:
            if char in s_statistics:
                s_statistics[char] += 1
            else:
                s_statistics[char] = 1
        t_statistics = {}
        for char in t:
            if char in t_statistics:
                t_statistics[char] += 1
            else:
                t_statistics[char] = 1
        for key in t_statistics.keys():
            if key not in s_statistics or s_statistics[key] != t_statistics[key]:
                return False
        return True
# @lc code=end

