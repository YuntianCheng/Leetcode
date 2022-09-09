#
# @lc app=leetcode.cn id=49 lang=python3
#
# [49] 字母异位词分组
#
import string
from typing import List
# @lc code=start
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        d = {}
        for str in strs:
            key = ''.join(sorted(str))
            if key in d:
                d[key].append(str)
            else:
                d[key] = [str]
        result = list(d.values())
        return result


# @lc code=end

