#
# @lc app=leetcode.cn id=539 lang=python3
#
# [539] 最小时间差
#
from math import sqrt
from typing import List
# @lc code=start
class Solution:
    def findMinDifference(self, timePoints: List[str]) -> int:
        time_in_min = []
        for t in timePoints:
            times = t.split(':')
            time_in_min.append(int(times[0])*60+int(times[1]))
        time_in_min.sort()
        min = 24*60
        for i in range(1,len(time_in_min)):
            current = self.duration(time_in_min[i], time_in_min[i-1])
            if current < min:
                min = current
        current = self.duration(time_in_min[len(time_in_min)-1], time_in_min[0])
        if current < min:
            min = current
        return min
    def duration(self, a:int, b:int) ->int:
        mid = 12*60
        d = abs(a-b)
        return d if d < mid else 24*60 - d
# @lc code=end

