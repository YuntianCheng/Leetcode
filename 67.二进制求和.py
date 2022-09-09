#
# @lc app=leetcode.cn id=67 lang=python3
#
# [67] 二进制求和
#

# @lc code=start
class Solution:
    def addBinary(self, a: str, b: str) -> str:
        up = 0
        a_l = len(a)
        b_l = len(b)
        result = list(a) if a_l > b_l else list(b)
        length = len(result)
        for i in range(1,length+1):
            first = '0'
            second = '0'
            if a_l - i >= 0:
                first = a[a_l - i]
            if b_l - i >= 0:
               second =  b[b_l-i]
            sum = int(first) + int(second)
            if up == 1:
                sum += 1
            if sum >= 2:
                result[length - i] = str(sum - 2)
                up = 1
            else:
                result[length - i] = str(sum)
                up = 0
        if up == 1:
            result = ['1'] + result
        return ''.join(result)


# @lc code=end

print(Solution().addBinary('1010','1011'))