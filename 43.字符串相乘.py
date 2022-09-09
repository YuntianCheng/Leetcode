#
# @lc app=leetcode.cn id=43 lang=python3
#
# [43] 字符串相乘
#

# @lc code=start
class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        num1_c = 1
        sum = 0
        for i in reversed(num1):
            num2_c = 1
            for j in reversed(num2):
                sum += int(i) * int(j) * num1_c * num2_c
                num2_c *= 10
            num1_c *= 10
        return str(sum)

# @lc code=end

print(Solution().multiply("123","456"))