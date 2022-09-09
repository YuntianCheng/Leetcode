#
# @lc app=leetcode.cn id=524 lang=python3
#
# [524] 通过删除字母匹配到字典里最长单词
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
    def findLongestWord(self, s: str, dictionary: List[str]) -> str:
        result = ''
        for i in range(len(dictionary)):
            if self.isSubsequence(dictionary[i], s):
                if len(dictionary[i]) > len(result) or (len(dictionary[i]) == len(result) and dictionary[i] < result):
                    result = dictionary[i]
        return result
# @lc code=end

