#
# @lc app=leetcode.cn id=125 lang=python3
#
# [125] 验证回文串
#

# @lc code=start
class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = s.lower()
        i = 0
        j = len(s) - 1
        if j == -1:
            return True
        while i < j:
            if (s[i] < 'a' or s[i] > 'z') and (s[i] < '0' or s[i] > '9'):
                i += 1
                continue
            if (s[j] < 'a' or s[j] > 'z') and (s[j] < '0' or s[j] > '9'):
                j -= 1
                continue
            if s[i] != s[j]:
                return False
            i += 1
            j -= 1
        return True
        
# @lc code=end

print(Solution().isPalindrome('1b1'))