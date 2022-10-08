//在给定的 m x n 网格
// grid 中，每个单元格可以有以下三个值之一：
//
//
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
//
//
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
//输出：4
//
//
// 示例 2：
//
//
//输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
//输出：-1
//解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
//
//
// 示例 3：
//
//
//输入：grid = [[0,2]]
//输出：0
//解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10
// grid[i][j] 仅为 0、1 或 2
//
//
// Related Topics 广度优先搜索 数组 矩阵 👍 609 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

func k1(grid [][]int, i, j, min int) {
	if grid[i][j] > 0 || grid[i][j] < -2-min {
		grid[i][j] = -2 - min
	}
	if i < len(grid)-1 {
		if grid[i+1][j] == 1 || (grid[i+1][j] != 0 && grid[i+1][j] < -min-2) {
			k(grid, i+1, j, min+1)
		}
	}
	if i > 0 {
		if grid[i-1][j] == 1 || (grid[i-1][j] != 0 && grid[i-1][j] < -min-2) {
			k(grid, i-1, j, min+1)
		}
	}
	if j < len(grid[i])-1 {
		if grid[i][j+1] == 1 || (grid[i][j+1] != 0 && grid[i][j+1] < -min-2) {
			k(grid, i, j+1, min+1)
		}
	}
	if j > 0 {
		if grid[i][j-1] == 1 || (grid[i][j-1] != 0 && grid[i][j-1] < -min-2) {
			k(grid, i, j-1, min+1)
		}
	}
}

func orangesRotting(grid [][]int) int {
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == 2 {
				k(grid, i, j, 0)
			}
		}
	}
	max := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
			if -grid[i][j] > max {
				max = -grid[i][j]
			}
		}
	}
	if max > 0 {
		return max - 2
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(orangesRotting(grid))
}
