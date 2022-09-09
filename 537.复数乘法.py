#
# @lc app=leetcode.cn id=537 lang=python3
#
# [537] 复数乘法
#

# @lc code=start
class Solution:
    def complexNumberMultiply(self, num1: str, num2: str) -> str:
        a1, b1 = map(int,num1[:-1].split('+'))
        a2,b2 = map(int,num2[:-1].split('+'))
        part_one = a1 * a2
        part_two = b1 * a2
        part_three = a1 * b2
        part_four =  0 - b1 * b2
        return '+'.join([str(part_one + part_four), str(part_three+part_two)+'i'])

# @lc code=end