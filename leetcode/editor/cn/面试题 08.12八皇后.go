//设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整
//个棋盘的那两条对角线。
//
// 注意：本题相对原题做了扩展
//
// 示例:
//
//  输入：4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释: 4 皇后问题存在如下两个不同的解法。
//[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//
//
// Related Topics 数组 回溯 👍 168 👎 0

package main

import (
	"fmt"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	used := make([]bool, n)
	first, second := make([]bool, 2*n), make([]bool, 2*n)
	var result [][]string
	var dfsol func(int, []string)
	dfsol = func(queensNum int, sol []string) {
		if queensNum == 0 {
			result = append(result, append([]string{}, sol...))
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] && !first[2*n-queensNum-i] && !second[n-queensNum+i] {
				var builder = strings.Builder{}
				for j := 0; j < i; j++ {
					builder.WriteByte('.')
				}
				builder.WriteByte('Q')
				for j := i + 1; j < n; j++ {
					builder.WriteByte('.')
				}
				used[i] = true
				first[2*n-queensNum-i] = true
				second[n-queensNum+i] = true
				sol = append(sol, builder.String())
				dfsol(queensNum-1, sol)
				used[i] = false
				first[2*n-queensNum-i] = false
				second[n-queensNum+i] = false
				sol = sol[:len(sol)-1]
			}
		}
	}
	dfsol(n, []string{})
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(solveNQueens(6))
}
