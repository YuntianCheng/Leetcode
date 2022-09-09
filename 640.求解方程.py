#
# @lc app=leetcode.cn id=640 lang=python3
#
# [640] 求解方程
#

# @lc code=start
import re


class Solution:
    def solveEquation(self, equation: str) -> str:
        left, right = equation.split('=')
        left_x_num = 0
        left_num = 0
        i = 0
        if left[0] != '-':
            left = ''.join(['+', left])
        for j in range(len(left)):
            if left[j] == '-' or left[j] == '+':
                if i != j:
                    mem = left[i:j]
                    if mem[-1] == 'x':
                        if len(mem) == 1:
                            left_x_num += 1
                        elif len(mem) == 2:
                            left_x_num += int(mem[0]+'1')
                        else:
                            left_x_num += int(mem[:-1])
                    else:
                        left_num += int(mem)
                    i = j
        mem = left[i:]
        if mem[-1] == 'x':
            if len(mem) == 1:
                left_x_num += 1
            elif len(mem) == 2:
                left_x_num += int(mem[0]+'1')
            else:
                left_x_num += int(mem[:-1])
        else:
            left_num += int(mem)
        right_x_num = 0
        right_num = 0
        i=0
        if right[0] != '-':
            right = ''.join(['+', right])
        for j in range(len(right)):
            if right[j] == '-' or right[j] == '+':
                if i != j:
                    mem = right[i:j]
                    if mem[-1] == 'x':
                        if len(mem) == 1:
                            right_x_num += 1
                        elif len(mem) == 2:
                            right_x_num += int(mem[0]+'1')
                        else:
                            right_x_num += int(mem[:-1])
                    else:
                        right_num += int(mem)
                    i = j
        mem = right[i:]
        if mem[-1] == 'x':
            if len(mem) == 1:
                right_x_num += 1
            elif len(mem) == 2:
                right_x_num += int(mem[0]+'1')
            else:
                right_x_num += int(mem[:-1])
        else:
            right_num += int(mem)
        x_num = left_x_num - right_x_num
        num = right_num - left_num
        if x_num == 0:
            if num == 0:
                return 'Infinite solutions'
            else:
                return 'No solution'
        result  = num/x_num
        if result-int(result) == 0:
            return 'x=' + str(int(result))
        return 'No solution'
        
        
                
# @lc code=end
print(Solution().solveEquation("x=x+2"))