#
# @lc app=leetcode.cn id=389 lang=python3
#
# [389] 找不同
#

# @lc code=start
class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        s_l = {}
        t_l = {}
        for char in s:
            if char in s_l:
                s_l[char] += 1
            else:
                s_l[char] = 1
        for char in t:
            if char in t_l:
                t_l[char] += 1
            else:
                t_l[char] = 1
        for key in t_l.keys():
            if key not in s_l or s_l[key] != t_l[key]:
                return key

# @lc code=end

