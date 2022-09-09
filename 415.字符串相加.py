#
# @lc app=leetcode.cn id=415 lang=python3
#
# [415] 字符串相加
#

# @lc code=start


class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        length = max(len(num1),len(num2))
        num1_len = len(num1)
        num2_len = len(num2)
        result = [''] * length
        up = 0
        for i in range(1, length + 1):
            first = 0
            second = 0
            if num1_len - i >= 0:
                first = int(num1[num1_len-i])
            if num2_len - i >= 0:
                second = int(num2[num2_len-i])
            sum = first + second + up
            if sum >= 10:
                result[length-i] = str(sum-10)
                up = 1
            else:
                result[length-i] = str(sum)
                up = 0
        if up == 1:
            result = ['1'] + result
        return ''.join(result)
# @lc code=end
print(Solution().addStrings('1','9'))

