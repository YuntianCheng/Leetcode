#
# @lc app=leetcode.cn id=520 lang=python3
#
# [520] 检测大写字母
#

# @lc code=start
a = 'a'
z = 'z'
A = 'A'
Z = 'Z'
class Solution:
    def detectCapitalUse(self, word: str) -> bool:
        hasLower = False
        hasUpper = False
        if word[0] >= a and word[0] <= z:
            hasLower = True
        for i in range(1,len(word), 1):
            if word[i] >= a and word[i] <= z:
                hasLower = True
            else:
                hasUpper = True
            if hasUpper and hasLower:
                return False
        return True

# @lc code=end

