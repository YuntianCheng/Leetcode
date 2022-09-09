#
# @lc app=leetcode.cn id=387 lang=python3
#
# [387] 字符串中的第一个唯一字符
#

# @lc code=start
class Solution:
    def firstUniqChar(self, s: str) -> int:
        letters = {}
        for char in s:
            if char in letters:
                letters[char]+=1
            else:
                letters[char] = 1
        for key in letters.keys():
            if letters[key] == 1:
                return s.index(key)
        return -1
# @lc code=end

