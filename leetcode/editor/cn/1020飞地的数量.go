//给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
//
// 一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
//
// 返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。
//
//
//
// 示例 1：
//
//
//输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
//输出：3
//解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
//
//
// 示例 2：
//
//
//输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
//输出：0
//解释：所有 1 都在边界上或可以到达边界。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 500
// grid[i][j] 的值为 0 或 1
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 186 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func edgeIsland(grid [][]int, i int, j int) (bool, int) {
	count := 0
	var self, left, right, top, bottom bool
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
		if grid[i][j] == 1 {
			count++
			grid[i][j] = -1
			if i == len(grid)-1 || i == 0 || j == len(grid[i])-1 || j == 0 {
				self = true
			}
			l, countL := edgeIsland(grid, i, j-1)
			r, countR := edgeIsland(grid, i, j+1)
			t, countT := edgeIsland(grid, i-1, j)
			b, countB := edgeIsland(grid, i+1, j)
			count += countB + countL + countR + countT
			left = l
			right = r
			top = t
			bottom = b
		}
	}
	if left || right || top || bottom || self {
		return true, count
	}
	return false, count
}
func numEnclaves(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				b, c := edgeIsland(grid, i, j)
				if !b {
					count += c
				}
			}
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(numEnclaves([][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}))
}
