#
# @lc app=leetcode.cn id=38 lang=python3
#
# [38] 外观数列
#

# @lc code=start
class Solution:
    def description(self, num:str)->str:
        result = []
        i = 0
        count = 1
        for j in range(1,len(num)):
            if num[j] == num[i]:
                count+=1
            else:
                result.append(str(count)+num[i])
                i = j
                count = 1
        result.append(str(count)+num[i])
        return ''.join(result)
    def countAndSay(self, n: int) -> str:
        no = '1'
        for i in range(1,n):
            no = self.description(no)
        return no
# @lc code=end

print(Solution().countAndSay(4))