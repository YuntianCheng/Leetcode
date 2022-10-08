//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例 1：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCCED"
//输出：true
//
//
// 示例 2：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"SEE"
//输出：true
//
//
// 示例 3：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCB"
//输出：false
//
//
//
//
// 提示：
//
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
//
//
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
//
// Related Topics 数组 回溯 矩阵 👍 1455 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func check(i, j, k int, word string, board [][]byte, visited [][]bool) bool {
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	visited[i][j] = true
	defer func() {
		visited[i][j] = false
	}()
	var l, r, t, b bool
	if i-1 >= 0 && !visited[i-1][j] {
		t = check(i-1, j, k+1, word, board, visited)
	}
	if i+1 < len(board) && !visited[i+1][j] {
		b = check(i+1, j, k+1, word, board, visited)
	}
	if j-1 >= 0 && !visited[i][j-1] {
		l = check(i, j-1, k+1, word, board, visited)
	}
	if j+1 < len(board[i]) && !visited[i][j+1] {
		r = check(i, j+1, k+1, word, board, visited)
	}
	return l || r || t || b
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if check(i, j, 0, word, board, visited) {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(exist([][]byte{
		{'a', 'b'},
	}, "ba"))
}
