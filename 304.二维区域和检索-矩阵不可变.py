#
# @lc app=leetcode.cn id=304 lang=python3
#
# [304] 二维区域和检索 - 矩阵不可变
#
from typing import List

# @lc code=start
class NumMatrix:

    def __init__(self, matrix: List[List[int]]):
        self.matrix = matrix
        for row in matrix:
            sum = 0
            for i in range(len(row)):
                sum = sum + row[i]
                row[i] = sum


    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        sum = 0
        for r in self.matrix[row1:row2 + 1]:
            sum  = sum + (r[col2] if col1 == 0 else (r[col2] - r[col1 - 1]))
        return sum




# Your NumMatrix object will be instantiated and called as such:
# obj = NumMatrix(matrix)
# param_1 = obj.sumRegion(row1,col1,row2,col2)
# @lc code=end

m = NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]])
m.sumRegion(2,1,4,3)