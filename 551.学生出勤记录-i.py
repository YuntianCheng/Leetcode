#
# @lc app=leetcode.cn id=551 lang=python3
#
# [551] 学生出勤记录 I
#

# @lc code=start


class Solution:
    def checkRecord(self, s: str) -> bool:
        countA = 0
        lhood = 0
        for c in s:
            if c == 'A':
                lhood = 0
                countA += 1
                if countA >= 2:
                    return False
            elif c == 'L':
                lhood += 1
                if lhood >= 3:
                    return False
            else:
                lhood = 0
        return True
# @lc code=end

