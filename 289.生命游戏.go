/*
 * @lc app=leetcode.cn id=289 lang=golang
 *
 * [289] 生命游戏
 */

// @lc code=start
func gameOfLife(board [][]int) {
	for i := range board {
		for j := range board[i] {
			result := IsAlive(board, i, j)
			if result != board[i][j] {
				if board[i][j] == 0 {
					board[i][j] = -1
				} else {
					board[i][j] = 2
				}
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == -1 {
				board[i][j] = 1
			} else if board[i][j] == 2 {
				board[i][j] = 0
			}
		}
	}
}
func IsAlive(board [][]int, row int, col int) int {
	m := len(board)
	n := len(board[0])
	sum := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i != row && i >= 0 && i < m && j != col && j >= 0 && j < n {
				if board[i][j] == 2 {
					sum += 1
				} else if board[i][j] == -1 {
					continue
				} else {
					sum += board[i][j]
				}
			}
		}
	}
	if sum == 3 {
		return 1
	} else if sum == 2 && (board[row][col] == 1 || board[row][col] == 2) {
		return 1
	}
	return 0
}

// @lc code=end

