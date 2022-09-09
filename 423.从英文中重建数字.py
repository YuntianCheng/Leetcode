#
# @lc app=leetcode.cn id=423 lang=python3
#
# [423] 从英文中重建数字
#

# @lc code=start
#one
#three
#five
#nine
chars = ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"]

class Solution:
    def originalDigits(self, s: str) -> str:
        result = [0]*10
        result[0] = s.count('z')
        result[2] = s.count('w')
        result[4] = s.count('u')
        result[8] = s.count('g')
        result[6] = s.count('x')
        result[7] = s.count('s') - result[6]
        result[5] = s.count('v') - result[7]
        return ''.join(sorted(result))
# @lc code=end

