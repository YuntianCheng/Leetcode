#
# @lc app=leetcode.cn id=592 lang=python3
#
# [592] 分数加减运算
#

# @lc code=start
import fractions
import re

class Solution:
    def fractionAddition(self, expression: str) -> str:
        nums = []
        i = 0
        k = 0
        if expression[0] == '-':
            k = 1
        for j in range(k,len(expression)):
            if expression[j] == '+' or expression[j] == '-':
                nums.append(expression[i:j])
                i = j
        nums.append(expression[i:])
        result = fractions.Fraction()
        for num in nums:
            fs = fractions.Fraction(num)
            result = result + fs
        return "/".join([str(result.numerator),str(result.denominator)])
# @lc code=end
print(Solution().fractionAddition("-5/2+10/3+7/9"))
