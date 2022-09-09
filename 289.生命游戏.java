/*
 * @lc app=leetcode.cn id=289 lang=java
 *
 * [289] 生命游戏
 */

// @lc code=start
class Solution {
    public void gameOfLife(int[][] board) {
        for(int i = 0; i < board.length; i++) {
            for(int j = 0; i < board[i].length; j++) {

            }
        }
    }

    public int IsAlive(int[][] board, int row, int col) {
        int m = board.length;
        int n = board[0].length;
        int num = 0
        for (int i = row - 1; i <= row + 1 && i != row; i++) {
            for(int j = col -1; j <= col + 1 && j != col; j++) {
                if (i >= 0 && i < m && j >=0 && j < n) {
                    num+=board[i][j]
                }
            }
        }
        if (num == 3) {
            return 1;
        }
        return 0;
    }
}
// @lc code=end

