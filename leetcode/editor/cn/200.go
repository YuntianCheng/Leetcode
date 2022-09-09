//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 1887 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

func addPoint(grid [][]byte, i int, j int) {
	if grid[i][j] == '1' {
		grid[i][j] = '2'
		if i > 0 {
			addPoint(grid, i-1, j)
		}
		if i+1 < len(grid) {
			addPoint(grid, i+1, j)
		}
		if j > 0 {
			addPoint(grid, i, j-1)
		}
		if j+1 < len(grid[i]) {
			addPoint(grid, i, j+1)
		}
	}
}

func numIslands(grid [][]byte) int {
	count := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == '1' {
				count++
				addPoint(grid, i, j)
			}
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}))
}
