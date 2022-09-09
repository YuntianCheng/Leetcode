#
# @lc app=leetcode.cn id=541 lang=python3
#
# [541] 反转字符串 II
#

# @lc code=start
class Solution:
    def reverseStr(self, s: str, k: int) -> str:
        sl = list(s)
        for i in range(0, len(s), 2*k):
            if i + k < len(s):
                m = i
                n = i + k - 1
                while m < n:
                    sl[m], sl[n] = sl[n], sl[m]
                    m+=1
                    n-=1
                continue
            m = i
            n = len(s) - 1
            while m < n:
                sl[m], sl[n] = sl[n], sl[m]
                m+=1
                n-=1
        return ''.join(sl)
            
        
# @lc code=end
print(Solution().reverseStr('abcdefg',2))
