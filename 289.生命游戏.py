#
# @lc app=leetcode.cn id=289 lang=python3
#
# [289] 生命游戏
#
from typing import List
# @lc code=start
class Solution:
    def isAlive(self, board: List[List[int]], row: int, col: int) -> int:
        m = len(board)
        n = len(board[0])
        sum = 0
        for i in range(row-1, row+2):
            if i >= 0 and i < m:
                for j in range(col-1, col+2):
                    if j == col and i == row:
                        continue
                    if j >= 0  and j < n:
                        if board[i][j] == 1 or board[i][j] == 3:
                            sum = sum + 1

        
        if sum == 3:
            return 1
        if sum == 2 and board[row][col] == 1:
            return 1
        return 0
    def gameOfLife(self, board: List[List[int]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        '''
        00,01,10,11
        '''
        for i in range(len(board)):
            for j in range(len(board[i])):
                life = self.isAlive(board, i, j)
                if life == 1:
                    board[i][j] = board[i][j] + 2
        for k in range(len(board)):
            for l in range(len(board[k])):
                if board[k][l] == 2 or board[k][l] == 3:
                    board[k][l] = 1
                else:
                    board[k][l] = 0

        


# @lc code=end

s = Solution()

board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
s.gameOfLife(board)