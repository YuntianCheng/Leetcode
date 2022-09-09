#
# @lc app=leetcode.cn id=467 lang=python3
#
# [467] 环绕字符串中唯一的子字符串
#

# @lc code=start


class Solution:
    def findSubstringInWraproundString(self, p: str) -> int:
        result = {}
        i = 0
        j = i + 1
        result[p[i]] = 1
        k = i
        while i < len(p) and j < len(p):
            if ord(p[j]) - ord(p[k]) == 1 or ord(p[j]) - ord(p[k]) == -25:
                if p[j] not in result or (p[j] in result and result[p[j]] < j - i + 1):
                    result[p[j]] = j- i + 1
                j += 1
                k += 1
            else:
                if p[j] not in result:
                    result[p[j]] = 1
                i = j
                j = i + 1
                k = i

        sum = 0
        for value in result.values():
            sum += value
        return sum



       
       
# @lc code=end
print(Solution().findSubstringInWraproundString('cdefghefghijklmnopqrstuvwxmnijklmnopqrstuvbcdefghijklmnopqrstuvwabcddefghijklfghijklmabcdefghijklmnopqrstuvwxymnopqrstuvwxyz'))
